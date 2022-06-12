package services

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

type UsersRepository interface {
	UserHash(ctx context.Context, login string) ([]byte, error)
	CreateUser(ctx context.Context, u *domain.User) error
}

type AuthService struct {
	usersRepo UsersRepository
}

func NewAuthService(repository UsersRepository) *AuthService {
	return &AuthService{
		usersRepo: repository,
	}
}

func (s *AuthService) Login(
	ctx context.Context,
	uc *domain.UserCredentials,
) error {
	actualHash, err := s.usersRepo.UserHash(ctx, uc.Login)
	if err != nil {
		return err
	}

	compareErr := bcrypt.CompareHashAndPassword(
		actualHash,
		[]byte(uc.Password),
	)
	if compareErr != nil {
		return domain.ErrAuthenticationFailed
	}

	return nil
}

func (s *AuthService) Register(
	ctx context.Context,
	uc *domain.UserCredentials,
) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(uc.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return s.usersRepo.CreateUser(ctx, &domain.User{
		Login: uc.Login,
		Hash:  hash,
	})
}
