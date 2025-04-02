package user_test

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/titanius-getha/gravitum-test-task/domain/user"
	"github.com/titanius-getha/gravitum-test-task/pkg/database"
)

func TestUserRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	pgContainer, err := postgres.Run(
		ctx,
		"postgres:16",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithLogger(log.New(io.Discard, "", 0)),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	assert.NoError(t, err)
	defer pgContainer.Terminate(ctx)

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.NewPostgres(dsn)
	assert.NoError(t, err)

	repo := user.NewPostgresRepository(db)
	err = repo.Migrate()
	assert.NoError(t, err)

	t.Run("Get unexist user", func(t *testing.T) {
		_, err := repo.GetByID(100000)
		assert.EqualError(t, err, user.ErrUserNotFound.Error(), "err should be user not found")
	})

	t.Run("Create user", func(t *testing.T) {
		user := user.User{Name: "User 1"}

		user, err := repo.Save(user)
		assert.NoError(t, err, "save user should have nil error")
		assert.NotZero(t, user.ID, "saved user id should not be zero")

		savedUser, err := repo.GetByID(user.ID)
		assert.NoError(t, err, "get by id after save should have nil error")
		assert.Equal(t, user, savedUser, "get by id user should be equal to saved user")
	})

	t.Run("Update user", func(t *testing.T) {
		user := user.User{Name: "User 2"}

		user, err := repo.Save(user)
		assert.NoError(t, err, "save user should have nil error")
		assert.NotZero(t, user.ID, "saved user id should not be zero")

		user.Name = "Updated user 2"
		updatedUser, err := repo.Save(user)
		assert.NoError(t, err, "save updated user should have nil error")
		assert.Equal(t, user, updatedUser, "updated user should be equal to saved user")

		takenUser, err := repo.GetByID(user.ID)
		assert.NoError(t, err, "get user after update should have nil error")
		assert.Equal(t, user, takenUser, "user taken after updated should be equal to updated user")
	})
}
