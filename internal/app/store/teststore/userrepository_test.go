package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x4AEKx/rest-api-go/internal/app/model"
	"github.com/x4AEKx/rest-api-go/internal/app/store"
	"github.com/x4AEKx/rest-api-go/internal/app/store/teststore"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
