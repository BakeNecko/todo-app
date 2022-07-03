package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-app/internal/entity"
)

type AuthRepository struct {
	collection *mongo.Collection
}

func NewAuthRepository(collection *mongo.Collection) *AuthRepository {
	return &AuthRepository{
		collection: collection,
	}
}

func (s *AuthRepository) Insert(ctx context.Context, dto *entity.AccountDTO) (string, error) {
	result, err := s.collection.InsertOne(ctx, dto)
	if err != nil {
		return "", fmt.Errorf("failter to create user due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (s *AuthRepository) GetByCredentials(ctx context.Context, email string, password string) (*entity.Account, error) {
	var a entity.Account
	res := s.collection.FindOne(
		ctx, bson.D{
			{"email", email},
			{"password", password},
		}, nil)
	if err := res.Err(); err != nil {
		return &a, err
	}
	if err := res.Decode(&a); err != nil {
		if err == mongo.ErrNoDocuments {
			return &a, errors.New("user doesn't exists")
		}
		return &a, err
	}
	return &a, nil

}

func (s *AuthRepository) SetSession(ctx context.Context, _id primitive.ObjectID, session *entity.Session) error {
	update := bson.M{"$set": bson.M{"session": session}}
	filter := bson.M{"_id": _id}
	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}
