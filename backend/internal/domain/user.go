package domain

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type User struct {
	ID        string
	Name      string
	Gender    Gender
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string, gender string) *User {
	return &User{
		ID:        ulid.Make().String(),
		Name:      name,
		Gender:    Gender(gender),
		CreatedAt: time.Now(),
	}
}

func ParseGender(g string) (Gender, error) {
	switch g {
	case "male", "Male":
		return GenderMale, nil
	case "female", "Female":
		return GenderFemale, nil
	default:
		return "", fmt.Errorf("invalid gender: %s", g)
	}
}
