package service

import (
	"apilaundry/config"
	"apilaundry/model"
	"apilaundry/model/dto"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(payload model.User) (dto.LoginResponseDto, error)
	VerifyToken(token string) (jwt.MapClaims, error)
}

type jwtService struct {
	config config.SecurityConfig
}

func (j *jwtService) GenerateToken(payload model.User) (dto.LoginResponseDto, error) {
	claims := dto.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.Durasi * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: payload.Id,
		Role:   payload.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.config.Key))
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	return dto.LoginResponseDto{Token: ss}, nil
}

func (j *jwtService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Key), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.config.Issuer {
		return nil, errors.New("invalid issuer or claims")
	}
	return claims, nil
}

func NewJwtService(cg config.SecurityConfig) JwtService {
	return &jwtService{config: cg}
}