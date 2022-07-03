package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-app/internal/entity"
)

// User for PSQL Logic
//type Account struct {
//	db    *sqlx.DB
//	table string
//}
//
//func NewAccount(db *sqlx.DB, table string) *Account {
//	return &Account{db, table}
//}
//
//func (s Account) Insert(ctx context.Context, account *core.Account) error {
//	// Some Code
//	return nil
//}

type accountRepository struct {
	collection *mongo.Collection
	//sLogger    *zap.SugaredLogger
}

func NewAccountRepository(collection *mongo.Collection) *accountRepository {
	//logger := logging.GetLogger()
	return &accountRepository{
		collection: collection,
		//sLogger:    logger.Sugar(),
	}
}

func (s accountRepository) GetByID(ctx context.Context, _id string) (*entity.Account, error) {
	var a *entity.Account

	objID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return a, err
	}

	res := s.collection.FindOne(ctx, bson.M{
		"_id": objID,
	}, nil)
	if err := res.Err(); err != nil {
		return a, err
	}
	if err := res.Decode(&a); err != nil {
		if err == mongo.ErrNoDocuments {
			return a, errors.New("user doesn't exists")
		}
		return a, err
	}
	return a, nil
}
