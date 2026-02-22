package dynamodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/mapper"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/model"
)

type TournamentRepository struct {
	*BaseRepository
}

func NewTournamentRepository(base *BaseRepository) *TournamentRepository {
	return &TournamentRepository{BaseRepository: base}
}

func (r *TournamentRepository) Create(ctx context.Context, t *domain.Tournament) error {
	item := mapper.ToTournamentItem(*t)

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = r.DB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           &r.cfg.Table,
		Item:                av,
		ConditionExpression: awsString("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TournamentRepository) CreateWithNewUsers(ctx context.Context, t *domain.Tournament, users []*domain.User) error {
	var transactItems []types.TransactWriteItem
	item := mapper.ToTournamentItem(*t)

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	transactItems = append(transactItems, types.TransactWriteItem{
		Put: &types.Put{
			TableName:           &r.cfg.Table,
			Item:                av,
			ConditionExpression: awsString("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
		},
	})

	for _, u := range users {
		userItem := mapper.ToUserItem(*u)
		uav, err := attributevalue.MarshalMap(userItem)
		if err != nil {
			return err
		}

		transactItems = append(transactItems, types.TransactWriteItem{
			Put: &types.Put{
				TableName:           &r.cfg.Table,
				Item:                uav,
				ConditionExpression: awsString("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
			},
		})
	}

	_, err = r.DB.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transactItems,
	})
	if err != nil {
		return fmt.Errorf("failed to create tournament and users: %w", err)
	}

	return nil
}

func (r *TournamentRepository) GetByID(ctx context.Context, tournamentID string) (*domain.Tournament, error) {
	out, err := r.DB.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.cfg.Table,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "TOURNAMENT#" + tournamentID},
			"SK": &types.AttributeValueMemberS{Value: "META"},
		},
	})
	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, nil
	}

	var item model.TournamentItem
	if err := attributevalue.UnmarshalMap(out.Item, &item); err != nil {
		return nil, err
	}

	t := mapper.ToDomainTournament(item)
	return t, nil
}

func (r *TournamentRepository) List(ctx context.Context) ([]*domain.Tournament, error) {

	out, err := r.DB.Query(ctx, &dynamodb.QueryInput{
		TableName:              &r.cfg.Table,
		IndexName:              awsString("GSI1"),
		KeyConditionExpression: awsString("GSI1PK = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "TOURNAMENT"},
		},
	})
	if err != nil {
		return nil, err
	}

	var items []model.TournamentItem
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &items); err != nil {
		return nil, err
	}

	result := make([]*domain.Tournament, 0, len(items))
	for _, it := range items {
		result = append(result, mapper.ToDomainTournament(it))
	}

	return result, nil
}

func (r *TournamentRepository) Update(ctx context.Context, tournamentID string, upd repository.TournamentUpdate) error {
	if upd.IsActive == nil && upd.Status == nil {
		return nil
	}

	updateParts := []string{}
	exprNames := map[string]string{}
	exprValues := map[string]types.AttributeValue{}

	if upd.IsActive != nil {
		exprNames["#is_active"] = "is_active"
		exprValues[":is_active"] = &types.AttributeValueMemberBOOL{Value: *upd.IsActive}
		updateParts = append(updateParts, "#is_active = :is_active")
	}

	if upd.Status != nil {
		exprNames["#status"] = "status"
		exprValues[":status"] = &types.AttributeValueMemberS{Value: string(*upd.Status)}
		updateParts = append(updateParts, "#status = :status")
	}

	exprNames["#updated_at"] = "updated_at"
	exprValues[":updated_at"] = &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)}
	updateParts = append(updateParts, "#updated_at = :updated_at")

	updateExpr := "SET " + strings.Join(updateParts, ", ")

	_, err := r.DB.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: awsString(r.cfg.Table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "TOURNAMENT#" + tournamentID},
			"SK": &types.AttributeValueMemberS{Value: "META"},
		},
		UpdateExpression:          awsString(updateExpr),
		ExpressionAttributeNames:  exprNames,
		ExpressionAttributeValues: exprValues,
		ConditionExpression:       awsString("attribute_exists(PK)"),
	})

	return err
}
