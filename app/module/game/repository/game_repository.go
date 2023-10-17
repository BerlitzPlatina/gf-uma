package repository

import (
	"context"

	"github.com/BerlitzPlatina/gf-uma/app/module/game/request"
	"github.com/BerlitzPlatina/gf-uma/internal/bootstrap/database"
	"github.com/BerlitzPlatina/gf-uma/internal/ent"
	"github.com/BerlitzPlatina/gf-uma/internal/ent/game"
)

type GameRepository struct {
	DB *database.Database
}

type IGameRepository interface {
	GetGames() ([]*ent.Game, error)
	GetGameByID(id int) (*ent.Game, error)
	CreateGame(request request.GameRequest) (*ent.Game, error)
	UpdateGame(id int, request request.GameRequest) (*ent.Game, error)
	DeleteGame(id int) error
}

func NewGameRepository(database *database.Database) *GameRepository {
	return &GameRepository{
		DB: database,
	}
}

func (s *GameRepository) GetGames() ([]*ent.Game, error) {
	return s.DB.Ent.Game.Query().Order(ent.Asc(game.FieldID)).All(context.Background())
}

func (s *GameRepository) GetGameByID(id int) (*ent.Game, error) {
	return s.DB.Ent.Game.Query().Where(game.IDEQ(id)).First(context.Background())
}

func (s *GameRepository) CreateGame(request request.GameRequest) (*ent.Game, error) {
	return s.DB.Ent.Game.Create().
		SetName(request.Name).
		SetIcon(request.Icon).
		Save(context.Background())
}

func (s *GameRepository) UpdateGame(id int, request request.GameRequest) (*ent.Game, error) {
	return s.DB.Ent.Game.UpdateOneID(id).
		SetName(request.Name).
		SetIcon(request.Icon).
		Save(context.Background())
}

func (s *GameRepository) DeleteGame(id int) error {
	return s.DB.Ent.Game.DeleteOneID(id).Exec(context.Background())
}
