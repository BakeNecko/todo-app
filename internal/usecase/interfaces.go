package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todo-app/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks/mocks.go -package=mocks

// UseCase Interfaces
type (
	IAccountUseCase interface {
		GetById(ctx context.Context, id string) (*entity.Account, error)
	}

	IAuthUseCase interface {
		SignIn(ctx context.Context, dto *entity.SignInDTO) (t Tokens, e error)
		SignUp(ctx context.Context, dto *entity.AccountDTO) (string, error)
	}

	ITodoUseCase interface {
		GetById(ctx context.Context, _id string) (*entity.TodoObject, error)
		Create(ctx context.Context, dto *entity.TodoDTO, accountID string) (string, error)
		GetAll(ctx context.Context, accountID string) ([]*entity.TodoObject, error)
		Update(ctx context.Context, dto *entity.TodoUpdateDTO, _id string) (*entity.TodoObject, error)
		Delete(ctx context.Context, _id string) error
	}
)

// Repository interfaces
type (
	IAuthRepository interface {
		Insert(ctx context.Context, dto *entity.AccountDTO) (string, error)
		GetByCredentials(ctx context.Context, email string, password string) (*entity.Account, error)
		SetSession(ctx context.Context, _id primitive.ObjectID, session *entity.Session) error
	}

	IAccountRepository interface {
		GetByID(ctx context.Context, _id string) (*entity.Account, error)
	}

	ITodoRepository interface {
		GetByID(ctx context.Context, todoId string) (*entity.TodoObject, error)
		GetAll(ctx context.Context, AccountId string) ([]*entity.TodoObject, error)
		Create(ctx context.Context, dto *entity.TodoDTO) (string, error)
		Update(ctx context.Context, dto *entity.TodoUpdateDTO, _id string) (string, error)
		Delete(ctx context.Context, todoId primitive.ObjectID) error
	}
)

// Other interfaces
type (
	// IPasswordHasher provides hashing logic to securely store passwords
	IPasswordHasher interface {
		Hash(password string) string
	}

	ITokenManager interface {
		NewJWT(_id string, ttl time.Duration) (string, error)
		Parse(accessToken string) (string, error)
		NewRefreshToken() (string, error)
	}
)

// Tokens represent jwt token pair
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
