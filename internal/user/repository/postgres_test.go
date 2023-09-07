package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aclgo/grpc-jwt/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Add(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewPostgresRepo(sqlxDB)

	columns := []string{"id", "name", "last_name", "password", "email", "role", "created_at", "updated_at"}
	uuidUser := uuid.NewString()
	now := time.Now()

	mockUser := &models.User{
		UserID:    uuidUser,
		Name:      "fake_name",
		Lastname:  "fake_lastname",
		Password:  "fake_pass",
		Email:     "email@gmail.com",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows(columns).AddRow(
		uuidUser,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	)

	mock.ExpectQuery(queryAddUser).WithArgs(
		mockUser.UserID,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	).WillReturnRows(rows)

	createdUser, err := userPGRepository.Add(context.Background(), mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
}

func TestFindByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepo := NewPostgresRepo(sqlxDB)

	columns := []string{"id", "name", "last_name", "password", "email", "role", "created_at", "updated_at"}
	uuidUser := uuid.NewString()
	now := time.Now()

	mockUser := models.User{
		UserID:    uuidUser,
		Name:      "fake_name",
		Lastname:  "fake_lastname",
		Password:  "fake_pass",
		Email:     "email@gmail.com",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.UserID,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	)

	mock.ExpectQuery(queryByID).WithArgs(mockUser.UserID).WillReturnRows(rows)

	foundUser, err := userPGRepo.FindByID(context.Background(), mockUser.UserID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.UserID, mockUser.UserID)
}

func TestFindByEmail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userRepo := NewPostgresRepo(sqlxDB)

	columns := []string{"id", "name", "last_name", "password", "email", "role", "created_at", "updated_at"}
	uuidUser := uuid.NewString()
	now := time.Now()

	mockUser := models.User{
		UserID:    uuidUser,
		Name:      "fake_name",
		Lastname:  "fake_lastname",
		Password:  "fake_pass",
		Email:     "email@gmail.com",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.UserID,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	)

	mock.ExpectQuery(queryFindByEmail).WithArgs(mockUser.Email).WillReturnRows(rows)

	foundUser, err := userRepo.FindByEmail(context.Background(), mockUser.Email)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.Email, mockUser.Email)
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userRepo := NewPostgresRepo(sqlxDB)

	columns := []string{"id", "name", "last_name", "password", "email", "role", "created_at", "updated_at"}
	uuidUser := uuid.NewString()
	now := time.Now()

	mockUser := models.User{
		UserID:    uuidUser,
		Name:      "fake_name",
		Lastname:  "fake_lastname",
		Password:  "fake_pass",
		Email:     "email@gmail.com",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.UserID,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	)

	mock.ExpectQuery(queryUpdate).WithArgs(
		mockUser.UserID,
		mockUser.Name,
		mockUser.Lastname,
		mockUser.Password,
		mockUser.Email,
		mockUser.Role,
		now,
		now,
	).WillReturnRows(rows)

	updatedUser, err := userRepo.Update(context.Background(), &mockUser)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	require.Equal(t, mockUser.UserID, uuidUser)

}
