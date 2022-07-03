package entity

import "go.mongodb.org/mongo-driver/bson"

type AccountDTO struct {
	Username string `json:"username" bson:"username" binding:"required,max=155"`
	Password string `json:"password" bson:"password" binding:"required,max=255"`
	Email    string `json:"email" bson:"email" binding:"required,max=155"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type TodoDTO struct {
	OwnerId string `json:"owner_id" bson:"owner_id"`
	Text    string `json:"text" bson:"text" binding:"required,max=1024"`
}

type TodoUpdateDTO struct {
	Text string `json:"text" bson:"text" binding:"max=1024"`
}

func (dto TodoUpdateDTO) Update() bson.M {
	update := bson.M{}

	if dto.Text != "" {
		update["text"] = dto.Text
	}
	return bson.M{"$set": update}
}

func (d *TodoDTO) SetOwnerId(OwnerId string) {
	d.OwnerId = OwnerId
}
