package service

import (
	"apilaundry/model"
	"apilaundry/model/dto"
	"apilaundry/repository"
	"apilaundry/util"
	"errors"
	"fmt"
)

type UserService interface {
	GetbyId(id string) (model.User, error)
	CreateNew(payload model.User) (model.User, error)
	Login(payload dto.LoginDto) (dto.LoginResponseDto, error)
	FindByUsername(username string) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
	jwt  JwtService
}

// GetById implements UserService
func (c *userService) GetbyId(id string) (model.User, error) {
	user, err := c.repo.GetbyId(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user %v tidak ditemukan!", id)
	}
	return user, nil
}

func (c *userService) CreateNew(payload model.User) (model.User, error) {
	if !payload.IsValidRole() {
		return model.User{}, errors.New("role is invalid, must be admin or employee")
	}
	passwordHash, err := util.EncryptPassword(payload.Password)
	fmt.Println("passwordhash:",passwordHash)
	if err != nil {
		return model.User{}, err
	}
	payload.Password = passwordHash
	return c.repo.CreateUser(payload)
}

func (c *userService) FindByUsername(username string) (model.User, error) {
	user, err := c.repo.FindByUsername(username)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (c *userService) Login(payload dto.LoginDto) (dto.LoginResponseDto, error) {
	user, err := c.repo.FindByUsername(payload.Username)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("username or password invalid ")
	}
	err = util.ComparePasswordHash(user.Password, payload.Password)
	if err != nil {
		fmt.Println(user.Password, payload.Password)
		return dto.LoginResponseDto{}, fmt.Errorf("password incorrect!")
	}
	user.Password = ""
	token, err := c.jwt.GenerateToken(user)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("failed to create token!")
	}
	return token, nil
}

func NewUserService(repositori repository.UserRepository, jwt JwtService) UserService {
	return &userService{repo: repositori, jwt: jwt}
}
