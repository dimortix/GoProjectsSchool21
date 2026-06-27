package datasource

import (
	"errors"
	"tictactoe/internal/domain"
)

type Repository interface {
	Save(game *domain.Game) error
	Get(uuid string) (*domain.Game, error)
}

type repositoryImpl struct {
	storage *Storage
}

func NewRepository(storage *Storage) Repository {
	return &repositoryImpl{
		storage: storage,
	}
}

func (r *repositoryImpl) Save(dGame *domain.Game) error {
	dsGame := toDsGame(dGame)
	r.storage.data.Store(dsGame.UUID, dsGame)
	return nil
}

func (r *repositoryImpl) Get(uuid string) (*domain.Game, error) {
	val, ok := r.storage.data.Load(uuid)
	if !ok {
		return nil, errors.New("game not found")
	}
	dsGame, ok := val.(*Game)
	if !ok {
		return nil, errors.New("invalid data in storage")
	}
	return toDomainGame(dsGame), nil
}
