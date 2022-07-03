package mongodb

import (
	"context"
	"errors"
	"fmt"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-app/internal/entity"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) *TodoRepository {
	return &TodoRepository{collection: collection}
}

func (s TodoRepository) GetByID(ctx context.Context, todoId string) (*entity.TodoObject, error) {
	var res *entity.TodoObject
	objId, err := entity.GetObjectID(todoId)
	if err != nil {
		return res, err
	}
	one := s.collection.FindOne(ctx, bson.M{"_id": objId})

	if err := one.Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return res, errors.New("object does not exists")
		}
		return res, err
	}
	return res, nil
}

func (s TodoRepository) GetAll(ctx context.Context, AccountId string) ([]*entity.TodoObject, error) {
	var res []*entity.TodoObject

	filter := bson.M{"owner_id": AccountId}
	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return res, err
	}
	if err = cursor.All(ctx, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (s TodoRepository) Create(ctx context.Context, dto *entity.TodoDTO) (string, error) {
	result, err := s.collection.InsertOne(ctx, dto)
	if err != nil {
		return "", fmt.Errorf("failter to create todo due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (s TodoRepository) Update(ctx context.Context, dto *entity.TodoUpdateDTO, _id string) (string, error) {
	oid, err := entity.GetObjectID(_id)
	if err != nil {
		return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s. due to error: %v", oid, err)
	}
	updated, err := s.collection.UpdateByID(ctx, oid, dto.Update())
	if err != nil {
		return "", err
	}
	if updated.MatchedCount == 0 {
		return "", fmt.Errorf("object not Found")
	}
	return _id, nil
}

func (s TodoRepository) Delete(ctx context.Context, todoId primitive.ObjectID) error {
	filter := bson.M{"_id": todoId}
	res, err := s.collection.DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		return fmt.Errorf("object(s) does not exists")
	}
	return err
}
