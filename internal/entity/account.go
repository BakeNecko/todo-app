package entity

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Session struct {
	RefreshToken string
	ExpiresAT    time.Time
}

type Account struct {
	ID           string       `json:"id" bson:"_id"`
	Username     string       `json:"username,omitempty" bson:"username"`
	Email        string       `json:"email,omitempty" bson:"email"`
	PasswordHash string       `json:"-" bson:"password_hash"`
	Session      *Session     `json:"-" bson:"session"`
	TodoList     []TodoObject `json:"todo_list" bson:"todo_list"`
}

type TodoObject struct {
	ID      string `json:"id" bson:"_id"`
	OwnerId string `json:"-" bson:"owner_id"`
	Text    string `json:"text" bson:"text"`
}

func GetObjectID(_id string) (oid primitive.ObjectID, e error) {
	hex, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		e = fmt.Errorf("can't cast current ID (%s) to primitive.GetObjectID. Error: %v", _id, err)
		return
	}
	return hex, nil
}
