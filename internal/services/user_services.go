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

func (s *UserService) UpdateUserDetails(userId string, req dto.UpdateUserProfile) (*model.User, error){
	user, err := s.GetById(userId)
	if err != nil{
		return nil, err
	}

	if req.Bio == nil &&
	req.GithubUsername == nil &&
	req.PortfolioWebsite == nil &&
	req.AvatarURL == nil {

	return nil, errors.New("at least one field must be provided")
    }
	

	if req.Bio != nil{
		if len(*req.Bio) > 500 {
			return nil, errors.New("bio cannot exceed 500 characters")
		}
		user.Bio = req.Bio
	}

	if req.GithubUsername != nil{
		if len(*req.GithubUsername) > 39 {
			return nil, errors.New("github username cannot exceed 39 characters")
		}
		user.GithubUsername = req.GithubUsername
	}

	if req.PortfolioWebsite != nil{
		if len(*req.PortfolioWebsite) > 255 {
			return nil, errors.New("portfolio website is too long")
		}

		if !utils.IsValidURL(*req.PortfolioWebsite){
			return nil, errors.New("invalid portfolio url")
		}
		user.PortfolioWebsite = req.PortfolioWebsite
	}
	

	if req.AvatarURL != nil{
		if len(*req.AvatarURL) > 255 {
			return nil, errors.New("avatar url is too long")
		}

		if !utils.IsValidURL(*req.AvatarURL) {
		return nil, errors.New("invalid avatar url")
		}
		user.AvatarURL= req.AvatarURL
	}

	return s.repo.UpdateUserDetails(user)
}

func (s *UserService) GetPublicProfile(username string) (*model.User, error){

	user, err := s.repo.GetByUsername(username)
	if err != nil{
		return nil, err
	}
	return user, nil
}