package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.UserRepo.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
