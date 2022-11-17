package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tomoropy/fishing-with-backend/util"
)

func createRandomUser(t *testing.T) Users {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Name:           util.RandomUserName(),
		ProfileText:    util.RandomUserProfileText(),
		Email:          util.RandomString(10),
		HashedPassword: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.ProfileText, user.ProfileText)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUer(t *testing.T) {
	user := createRandomUser(t)
	GetedUser, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, GetedUser)

	require.Equal(t, user.ID, GetedUser.ID)
	require.Equal(t, user.Name, GetedUser.Name)
	require.Equal(t, user.ProfileText, GetedUser.ProfileText)
	require.Equal(t, user.Email, GetedUser.Email)
	require.Equal(t, user.HashedPassword, GetedUser.HashedPassword)
	require.WithinDuration(t, user.CreatedAt.Time, GetedUser.CreatedAt.Time, time.Second)

}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	arg := UpdateUserParams{
		ID:   user.ID,
		Name: util.RandomString(10),
		ProfileText: util.RandomString(100),
		Email: util.RandomEmail(),
		HashedPassword: util.RandomString(100),
	}

	err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	deletedUser, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedUser)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUserParams{
		Limit:  5,
		Offset: 0,
	}
	listUser, err := testQueries.ListUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listUser, 5)

	for _, user := range listUser {
		require.NotEmpty(t, user)
	}

}
