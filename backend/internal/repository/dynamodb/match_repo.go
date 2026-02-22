package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/mapper"
)

type MatchRepository struct {
	*BaseRepository
}

func NewMatchRepository(base *BaseRepository) *MatchRepository {
	return &MatchRepository{BaseRepository: base}
}

func (r *MatchRepository) BatchCreate(ctx context.Context, matches []*domain.Match) error {
	batchSize := r.cfg.BatchSize
	if batchSize <= 0 || batchSize > 25 {
		batchSize = 25
	}

	for i := 0; i < len(matches); i += batchSize {
		end := i + batchSize
		if end > len(matches) {
			end = len(matches)
		}
		batch := matches[i:end]

		if err := r.writeBatchWithRetry(ctx, batch); err != nil {
			return err
		}
	}

	return nil
}

func (r *MatchRepository) writeBatchWithRetry(ctx context.Context, matches []*domain.Match) error {
	writeReqs := make([]types.WriteRequest, 0, len(matches))
	for _, m := range matches {
		item := mapper.ToMatchItem(*m)
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			return err
		}
		writeReqs = append(writeReqs, types.WriteRequest{
			PutRequest: &types.PutRequest{Item: av},
		})
	}

	requestItems := map[string][]types.WriteRequest{
		r.cfg.Table: writeReqs,
	}

	backoff := r.cfg.RetryBackoff
	if backoff <= 0 {
		backoff = 100 * time.Millisecond
	}
	maxRetries := r.cfg.NumRetries
	if maxRetries <= 0 {
		maxRetries = 5
	}

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		out, err := r.DB.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: requestItems,
		})
		if err != nil {
			lastErr = err
			if attempt < maxRetries {
				time.Sleep(backoff * time.Duration(attempt+1))
			}
			continue
		}

		if len(out.UnprocessedItems) == 0 {
			return nil
		}
		requestItems = out.UnprocessedItems
		if attempt < maxRetries {
			time.Sleep(backoff * time.Duration(attempt+1))
		} else {
			// Exhausted retries with unprocessed items remaining
			totalUnprocessed := 0
			for _, reqs := range requestItems {
				totalUnprocessed += len(reqs)
			}
			return fmt.Errorf("batch write failed: %d item(s) unprocessed after %d retries (possible throttling)", totalUnprocessed, maxRetries+1)
		}
	}

	return lastErr
}

func (r *MatchRepository) GetByID(ctx context.Context, matchID string) (*domain.Match, error) {
	return nil, nil
}

func (r *MatchRepository) SubmitScore(ctx context.Context, match domain.Match) error {
	return nil
}

func (r *MatchRepository) ListByTournamentID(ctx context.Context, tournamentID string) ([]*domain.Match, error) {
	return nil, nil
}
