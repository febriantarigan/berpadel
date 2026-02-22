package mapper

import (
	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/model"
	"github.com/febriantarigan/berpadel/internal/util"
)

func ToUserItem(u domain.User) *model.UserItem {
	return &model.UserItem{
		PK:        "USER#" + u.ID,
		SK:        "PROFILE",
		Name:      u.Name,
		Gender:    string(u.Gender),
		CreatedAt: util.ToString(u.CreatedAt),
		UpdatedAt: util.ToString(u.UpdatedAt),
	}
}

func ToDomainUser(item model.UserItem) *domain.User {
	return &domain.User{
		ID:        extractKey("USER#", item.PK),
		Name:      item.Name,
		Gender:    domain.Gender(item.Gender),
		CreatedAt: util.ToTime(item.CreatedAt),
		UpdatedAt: util.ToTime(item.UpdatedAt),
	}
}
