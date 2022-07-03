package usecase

import (
	"time"
)

type Deps struct {
	AccountRepository IAccountRepository
	AuthRepository    IAuthRepository
	TodoRepository    ITodoRepository

	Hasher          IPasswordHasher
	JwtManager      ITokenManager
	RefreshTokenTTL time.Duration
	AccessTokenTTL  time.Duration
}

type UseCases struct {
	AuthUseCase    IAuthUseCase
	AccountUseCase IAccountUseCase
	TodoUseCase    ITodoUseCase
}

func NewUseCases(deps *Deps) *UseCases {
	return &UseCases{
		AuthUseCase: NewAuthUseCase(
			deps.AuthRepository,
			deps.Hasher,
			deps.JwtManager,
			deps.RefreshTokenTTL,
			deps.AccessTokenTTL,
		),
		AccountUseCase: NewAccountUseCase(
			deps.AccountRepository,
			deps.Hasher,
		),
		TodoUseCase: NewTodoUseCase(
			deps.TodoRepository,
		),
	}
}
