package service

import (
	"errors"
	"fmt"
	"github.com/rizdian/go-full-api/internal/model"
	"github.com/rizdian/go-full-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserService interface {
	Create(user *model.User) error
	GetByID(id string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(user *model.User) error {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Save user to DB
	return s.userRepo.Create(user)
}

func (s *userService) GetByID(id string) (*model.User, error) {
	// Convert id to uint and handle error
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	// Convert to uint as GORM expects uint
	user, err := s.userRepo.FindByID(uint(userID))
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
