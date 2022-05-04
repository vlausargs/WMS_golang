package db

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	"github.com/hidratechh/GO_WMS/util"

// 	"github.com/stretchr/testify/require"

// 	sqlc "github.com/hidratechh/GO_WMS/db/sqlc"
// )

// func createRandomUser(t *testing.T) sqlc.User {
// 	arg := sqlc.CreateUserParams{
// 		FullName: sql.NullString{String: util.RandomName(), Valid: true},
// 		Username: sql.NullString{String: util.RandomName(), Valid: true},
// 		Password: sql.NullString{String: util.RandomName(), Valid: true},
// 		Photo:    sql.NullString{String: "/var/project/noqueue/img.jpg", Valid: true},
// 	}
// 	user, err := testQuery.CreateUser(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)

// 	require.Equal(t, arg.FullName, user.FullName)
// 	require.Equal(t, arg.Username, user.Username)
// 	require.Equal(t, arg.Password, user.Password)
// 	require.Equal(t, arg.Photo, user.Photo)
// 	require.NotZero(t, user.ID)

// 	require.NotZero(t, user.CreatedAt)
// 	return user
// }
// func TestCreateUser(t *testing.T) {
// 	createRandomUser(t)
// }

// func TestGetUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testQuery.GetUser(context.Background(), user1.ID)
// 	require.NoError(t, err)

// 	require.NotEmpty(t, user2)
// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.FullName, user2.FullName)
// 	require.Equal(t, user1.Username, user2.Username)
// 	require.Equal(t, user1.Password, user2.Password)
// 	require.Equal(t, user1.Photo, user2.Photo)

// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }
