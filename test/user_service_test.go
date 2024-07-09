package test

import (
	"context"
	"testing"

	"github.com/Pawancod/gRPC-User-Service/internal/repository"
	"github.com/Pawancod/gRPC-User-Service/internal/service"
	"github.com/Pawancod/gRPC-User-Service/proto/github.com/Pawancod/gRPC-User-Service/proto"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)

	req := &proto.GetUserRequest{Id: 1}
	res, err := userService.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Id, res.User.Id)
}

func TestGetUsers(t *testing.T) {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)

	req := &proto.GetUsersRequest{Ids: []int32{1, 2}}
	res, err := userService.GetUsers(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res.Users))
}

func TestSearchUsers(t *testing.T) {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)

	req := &proto.SearchUsersRequest{City: "LA"}
	res, err := userService.SearchUsers(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res.Users) > 0)
}
