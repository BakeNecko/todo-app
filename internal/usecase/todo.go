package usecase

import (
	"context"
	"todo-app/internal/entity"
)

type TodoUseCase struct {
	storage ITodoRepository
}

func (s TodoUseCase) GetById(ctx context.Context, _id string) (*entity.TodoObject, error) {
	return s.storage.GetByID(ctx, _id)
}

func (s TodoUseCase) Create(ctx context.Context, dto *entity.TodoDTO, accountID string) (string, error) {
	dto.SetOwnerId(accountID)
	return s.storage.Create(ctx, dto)
}

func (s TodoUseCase) Update(ctx context.Context, dto *entity.TodoUpdateDTO, _id string) (*entity.TodoObject, error) {
	var res *entity.TodoObject
	_id, err := s.storage.Update(ctx, dto, _id)
	if err != nil {
		return res, err
	}
	return s.storage.GetByID(ctx, _id)
}

func (s TodoUseCase) GetAll(ctx context.Context, accountID string) ([]*entity.TodoObject, error) {
	return s.storage.GetAll(ctx, accountID)
}

func (s TodoUseCase) Delete(ctx context.Context, _id string) error {
	objId, err := entity.GetObjectID(_id)
	if err != nil {
		return err
	}
	return s.storage.Delete(ctx, objId)
}

func NewTodoUseCase(repository ITodoRepository) *TodoUseCase {
	return &TodoUseCase{
		storage: repository,
	}
}
