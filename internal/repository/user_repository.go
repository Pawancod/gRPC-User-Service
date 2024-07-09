package repository

import (
    "errors"
    "github.com/Pawancod/gRPC-User-Service/internal/model"
)

var users = []model.User{
    {ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
    {ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 5.9, Married: false},
    // Add more users as needed
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) GetUserByID(id int32) (model.User, error) {
    for _, user := range users {
        if user.ID == id {
            return user, nil
        }
    }
    return model.User{}, errors.New("user not found")
}

func (r *UserRepository) GetUsersByID(ids []int32) ([]model.User, error) {
    var result []model.User
    for _, id := range ids {
        for _, user := range users {
            if user.ID == id {
                result = append(result, user)
            }
        }
    }
    return result, nil
}

func (r *UserRepository) SearchUsers(city string, phone int64, married bool) ([]model.User, error) {
    var result []model.User
    for _, user := range users {
        if (city == "" || user.City == city) &&
           (phone == 0 || user.Phone == phone) &&
           (married == false || user.Married == married) {
            result = append(result, user)
        }
    }
    return result, nil
}
