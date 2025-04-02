package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/titanius-getha/gravitum-test-task/domain/user"
	mocks "github.com/titanius-getha/gravitum-test-task/mocks/github.com/titanius-getha/gravitum-test-task/domain/user"
)

func TestUserService_GetUserByID(t *testing.T) {
	repo := new(mocks.MockUserRepository)

	expectedUser := user.User{ID: 1, Name: "User 1"}

	repo.EXPECT().GetByID(uint(1)).Return(expectedUser, nil)
	repo.EXPECT().Migrate().Return(nil)

	service, err := user.NewService(repo)
	assert.NoError(t, err, "migration error should be nil")

	user, err := service.GetByID(expectedUser.ID)
	assert.NoError(t, err, "get user by id error should be nil")
	assert.Equal(t, expectedUser, user, "user should be equal to expected user")
}

func TestUserService_GetUnexistUserById(t *testing.T) {
	repo := new(mocks.MockUserRepository)

	repo.On("GetByID", mock.Anything).Return(user.User{}, user.ErrUserNotFound)
	repo.EXPECT().Migrate().Return(nil)

	service, err := user.NewService(repo)
	assert.NoError(t, err, "migration error should be nil")

	_, err = service.GetByID(1)
	assert.EqualError(t, err, user.ErrUserNotFound.Error())
}

func TestUserService_CreateUser(t *testing.T) {
	repo := new(mocks.MockUserRepository)

	expectedUser := user.User{Name: "User 1"}

	repo.EXPECT().Save(expectedUser).Return(expectedUser, nil)
	repo.EXPECT().Migrate().Return(nil)

	service, err := user.NewService(repo)
	assert.NoError(t, err, "migration error should be nil")

	createdUser, err := service.Create(expectedUser.Name)
	assert.NoError(t, err, "create user error should be nil")
	assert.Equal(t, expectedUser.Name, createdUser.Name, "created user should be equal to expected user")
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := new(mocks.MockUserRepository)

	originalUser := user.User{ID: 1, Name: "User 1"}
	expectedUser := user.User{ID: 1, Name: "User 2"}

	repo.EXPECT().GetByID(uint(originalUser.ID)).Return(originalUser, nil)
	repo.EXPECT().Save(expectedUser).Return(expectedUser, nil)
	repo.EXPECT().Migrate().Return(nil)

	service, err := user.NewService(repo)
	assert.NoError(t, err, "migration error should be nil")

	updatedUser, err := service.Update(originalUser.ID, expectedUser.Name)
	assert.NoError(t, err, "update user error should be nil")
	assert.Equal(t, expectedUser, updatedUser, "updated user should be equal to expected user")
}

func TestUserService_UpdateUnexistUser(t *testing.T) {
	repo := new(mocks.MockUserRepository)

	repo.On("GetByID", mock.Anything).Return(user.User{}, user.ErrUserNotFound)
	repo.EXPECT().Migrate().Return(nil)

	service, err := user.NewService(repo)
	assert.NoError(t, err, "migration error should be nil")

	_, err = service.Update(1, "New name")
	assert.EqualError(t, err, user.ErrUserNotFound.Error(), "update unexist error should be equal to ErrUserNotFound")

	repo.AssertNotCalled(t, "Save")
}
