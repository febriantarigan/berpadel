package util

import (
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
)

func ToString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func ToTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func UsersToStrings(users []*domain.User) []string {
	userIDs := make([]string, 0, len(users))
	for _, u := range users {
		userIDs = append(userIDs, u.ID)
	}
	return userIDs
}

func StringsToUsers(userIDs []string) []*domain.User {
	users := make([]*domain.User, 0, len(userIDs))
	for _, u := range userIDs {
		users = append(users, &domain.User{ID: u})
	}
	return users
}
