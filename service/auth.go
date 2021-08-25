package service

import (
	"context"
	"go-prisma/prisma/db"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	client db.PrismaClient
}

func NewAuthService(client db.PrismaClient) *AuthService {
	return &AuthService{client}
}

func (s *AuthService) ValidateUser(email string, password string) (*db.UserModel, error) {
	user, err := s.client.User.FindFirst(db.User.Email.Equals(email)).Exec(context.Background())
	if err != nil {
		return &db.UserModel{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return &db.UserModel{}, err
	}

	return user, nil
}
