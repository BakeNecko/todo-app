package usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-app/internal/entity"
)

type AccountUseCase struct {
	repository IAccountRepository
	hasher     IPasswordHasher
}

func NewAccountUseCase(repository IAccountRepository, hasher IPasswordHasher) *AccountUseCase {
	return &AccountUseCase{
		repository: repository,
		hasher:     hasher,
	}
}

func (a *AccountUseCase) GetById(ctx context.Context, id string) (*entity.Account, error) {
	account, err := a.repository.GetByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return account, nil
}
