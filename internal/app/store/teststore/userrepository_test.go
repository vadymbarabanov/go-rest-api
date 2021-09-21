package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vadymbarabanov/go-rest-api/internal/app/model"
	"github.com/vadymbarabanov/go-rest-api/internal/app/store"
	"github.com/vadymbarabanov/go-rest-api/internal/app/store/teststore"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	if err := s.User().Create(u); err != nil {
		t.Fatal(err)
	}

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
