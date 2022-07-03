package collections

import "go.mongodb.org/mongo-driver/mongo"

const (
	accountCollection = "accounts"
	toDoCollection    = "todo"
)

type Collections struct {
	Account *mongo.Collection
	ToDo    *mongo.Collection
}

func New(database *mongo.Database) *Collections {
	return &Collections{
		Account: database.Collection(accountCollection),
		ToDo:    database.Collection(toDoCollection),
	}
}
