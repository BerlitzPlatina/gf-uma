package service

import (
	"github.com/BerlitzPlatina/gf-uma/app/module/game/repository"
	"github.com/BerlitzPlatina/gf-uma/app/module/game/request"
	"github.com/BerlitzPlatina/gf-uma/internal/ent"
)

type GameService struct {
	Repo *repository.GameRepository
}

type IGameService interface {
	GetGames() ([]*ent.Game, error)
	GetGameByID(id int) (*ent.Game, error)
	CreateGame(request request.GameRequest) (*ent.Game, error)
	UpdateGame(id int, request request.GameRequest) (*ent.Game, error)
	DeleteGame(id int) error
}

func NewGameService(repo *repository.GameRepository) *GameService {
	return &GameService{
		Repo: repo,
	}
}

func (s *GameService) GetGames() ([]*ent.Game, error) {
	return s.Repo.GetGames()
}

func (s *GameService) GetGameByID(id int) (*ent.Game, error) {
	return s.Repo.GetGameByID(id)
}

func (s *GameService) CreateGame(request request.GameRequest) (*ent.Game, error) {
	return s.Repo.CreateGame(request)
}

func (s *GameService) UpdateGame(id int, request request.GameRequest) (*ent.Game, error) {
	return s.Repo.UpdateGame(id, request)
}

func (s *GameService) DeleteGame(id int) error {
	return s.Repo.DeleteGame(id)
}
