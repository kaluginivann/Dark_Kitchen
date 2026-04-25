package users

import (
	"context"

	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Register(ctx context.Context, payload *RegisterRequest) (*RegisterResponse, error)
}

type UserService struct {
	UserRepository IRepository
	Logger         logger.Interface
}

func NewService(userRepository IRepository, logger logger.Interface) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Logger:         logger,
	}
}

func (u *UserService) Register(ctx context.Context, payload *RegisterRequest) (*RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	payload.Password = string(hashedPassword)
	obj, err := u.UserRepository.CreateUser(ctx, payload.Username, payload.Email, payload.Password)
	if err != nil {
		return nil, err
	}
	var registerResp RegisterResponse

	registerResp.Id = obj.Id
	registerResp.Username = obj.Username
	registerResp.Email = obj.Email

	return &registerResp, nil
}
