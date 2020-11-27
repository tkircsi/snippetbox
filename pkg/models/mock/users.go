package mock

import (
	"time"

	"github.com/tkircsi/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Boldi",
	Email:   "boldi@gmail.com",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case mockUser.Email:
		return mockUser.ID, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case mockUser.ID:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
