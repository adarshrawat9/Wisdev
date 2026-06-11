package services

import (
	"Wisdev/internal/dto"
	"Wisdev/internal/model"
	"Wisdev/internal/repositories"
	"Wisdev/internal/utils"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UserService struct{
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}


func (s *UserService) Register(req dto.RegisterRequest)(*model.User, error){

	user, err := s.repo.GetByEmail(req.Email)
	if err == nil{
		return nil, errors.New("email already exist")
	}
	if err != pgx.ErrNoRows{
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil{
		return nil, err
	}

	user = &model.User{
		Email : req.Email,
		Username: req.Username,
		PasswordHash: hashedPassword,
		
	}

	err = s.repo.CreateUser(user)

	if err != nil{
		return nil, err
	}

	return user, nil

}


func (s *UserService) Login(req dto.LoginRequest) (string, error){

	user, err := s.repo.GetByEmail(req.Email)
	if err == pgx.ErrNoRows{
		return "", errors.New("Invalid credentials")
	}
	if err != nil{
		return "", err
	}

    err = utils.VerifyPassword(user.PasswordHash, req.Password)
	if err != nil{
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil{
		return "", err
	}

	return token, nil
	
}

func (s *UserService) GetById(userId string) (*model.User, error){
	user, err := s.repo.GetById(userId)
	if err != nil{
		return nil, err
	}
	return user, nil
}
