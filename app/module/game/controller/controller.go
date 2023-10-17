package controller

import "github.com/efectn/fiber-boilerplate/app/module/game/service"

type Controller struct {
	Game *GameController
}

func NewController(gameService *service.GameService) *Controller {
	return &Controller{
		Game: &GameController{gameService: gameService},
	}
}
