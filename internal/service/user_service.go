package service

import (
	"context"

	"github.com/Pawancod/gRPC-User-Service/internal/repository"
	"github.com/Pawancod/gRPC-User-Service/proto/github.com/Pawancod/gRPC-User-Service/proto"
)

type UserService struct {
    proto.UnimplementedUserServiceServer
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
    user, err := s.repo.GetUserByID(req.Id)
    if err != nil {
        return nil, err
    }
    return &proto.GetUserResponse{
        User: &proto.User{
            Id:      user.ID,
            Fname:   user.FName,
            City:    user.City,
            Phone:   user.Phone,
            Height:  user.Height,
            Married: user.Married,
        },
    }, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
    users, err := s.repo.GetUsersByID(req.Ids)
    if err != nil {
        return nil, err
    }
    var protoUsers []*proto.User
    for _, user := range users {
        protoUsers = append(protoUsers, &proto.User{
            Id:      user.ID,
            Fname:   user.FName,
            City:    user.City,
            Phone:   user.Phone,
            Height:  user.Height,
            Married: user.Married,
        })
    }
    return &proto.GetUsersResponse{Users: protoUsers}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, req *proto.SearchUsersRequest) (*proto.SearchUsersResponse, error) {
    users, err := s.repo.SearchUsers(req.City, req.Phone, req.Married)
    if err != nil {
        return nil, err
    }
    var protoUsers []*proto.User
    for _, user := range users {
        protoUsers = append(protoUsers, &proto.User{
            Id:      user.ID,
            Fname:   user.FName,
            City:    user.City,
            Phone:   user.Phone,
            Height:  user.Height,
            Married: user.Married,
        })
    }
    return &proto.SearchUsersResponse{Users: protoUsers}, nil
}
