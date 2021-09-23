package store

import "github.com/vadymbarabanov/go-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByID(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
