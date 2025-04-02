package service

import (
	"errors"
	"techdocs/internal/model"
	"techdocs/internal/repository"
	"techdocs/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

func (s *UserService) Register(req *RegisterRequest) error {
	// Check if username exists
	_, err := s.userRepo.FindByUsername(req.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	// Check if email exists
	_, err = s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     "user",
	}

	return s.userRepo.Create(user)
}

func (s *UserService) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	// Clear password from response
	user.Password = ""

	return &LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Password = "" // Clear password from response
	return user, nil
}

func (s *UserService) UpdateUser(id uint, updates *model.User) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Only allow updating certain fields
	user.Email = updates.Email
	if updates.Password != "" {
		user.Password = updates.Password
	}

	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) ListUsers(page, limit int) ([]model.User, int64, error) {
	users, total, err := s.userRepo.List(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Clear passwords from response
	for i := range users {
		users[i].Password = ""
	}

	return users, total, nil
}

func (s *UserService) GetJWTSecret() string {
	return s.jwtSecret
}
