package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"todo-app/internal/usecase"
	"todo-app/internal/usecase/repository/mongodb"
	"todo-app/internal/usecase/repository/mongodb/collections"
)

type Repository struct {
	AuthRepository    usecase.IAuthRepository
	AccountRepository usecase.IAccountRepository
	TodoRepository    usecase.ITodoRepository
}

func NewRepository(mdb *mongo.Database) *Repository {
	colls := collections.New(mdb)
	return &Repository{
		AuthRepository:    mongodb.NewAuthRepository(colls.Account),
		AccountRepository: mongodb.NewAccountRepository(colls.Account),
		TodoRepository:    mongodb.NewTodoRepository(colls.Account),
	}
}
