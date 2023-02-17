package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createNewUser(t)
}

func TestGetUser(t *testing.T) {
	user := createNewUser(t)
	getUser, err := testRepo.GetUser(context.Background(), testDB, user.ID)
	require.NoError(t, err)
	require.Equal(t, user, getUser)
}

func TestGetBatchUsers(t *testing.T) {
	count := 5
	var ids []int64
	for i := 0; i < count; i++ {
		user := createNewUser(t)
		ids = append(ids, user.ID)
	}

	results, err := testRepo.GetBatchUsers(context.Background(), testDB, ids)
	require.NoError(t, err)
	require.Equal(t, len(results), count)

	fmt.Println("ids", ids)
	users := make(map[int64]User, len(results))
	for _, result := range results {
		fmt.Println("user id", result.ID)
		users[result.ID] = result
	}

	// testing
	for _, id := range ids {
		user, found := users[id]
		require.True(t, found)
		require.Equal(t, id, user.ID)
	}
}

func createNewUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:  "royyan",
		Email: "royyan@gmail.com",
	}

	user, err := testRepo.CreateUser(context.Background(), testDB, arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}
