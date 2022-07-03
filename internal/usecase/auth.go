package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todo-app/internal/entity"
)

type AuthUseCase struct {
	repository      IAuthRepository
	hasher          IPasswordHasher
	tokenManager    ITokenManager
	refreshTokenTTL time.Duration
	accessTokenTTL  time.Duration
}

func NewAuthUseCase(
	storage IAuthRepository,
	hasher IPasswordHasher,
	manger ITokenManager,
	refreshTokenTTL time.Duration,
	accessTokenTTL time.Duration,
) *AuthUseCase {
	return &AuthUseCase{
		repository:      storage,
		hasher:          hasher,
		tokenManager:    manger,
		refreshTokenTTL: refreshTokenTTL,
		accessTokenTTL:  accessTokenTTL,
	}
}

func (s *AuthUseCase) SignUp(ctx context.Context, dto *entity.AccountDTO) (string, error) {
	dto.Password = s.hasher.Hash(dto.Password)
	_id, err := s.repository.Insert(ctx, dto)
	if err != nil {
		return "", err
	}
	return _id, nil
}

func (s *AuthUseCase) SignIn(ctx context.Context, dto *entity.SignInDTO) (t Tokens, e error) {
	hashPassword := s.hasher.Hash(dto.Password)
	account, err := s.repository.GetByCredentials(ctx, dto.Email, hashPassword)
	if err != nil {
		return Tokens{}, err
	}
	oid, err := entity.GetObjectID(account.ID)
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(ctx, oid)
}

func (s *AuthUseCase) createSession(ctx context.Context, _id primitive.ObjectID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(_id.Hex(), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := &entity.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAT:    time.Now().Add(s.refreshTokenTTL),
	}
	err = s.repository.SetSession(ctx, _id, session)
	return res, err
}
