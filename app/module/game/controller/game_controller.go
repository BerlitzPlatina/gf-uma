package controller

import (
	"strconv"

	"github.com/BerlitzPlatina/gf-uma/app/module/game/request"
	"github.com/BerlitzPlatina/gf-uma/app/module/game/service"
	"github.com/BerlitzPlatina/gf-uma/utils/response"
	"github.com/gofiber/fiber/v2"
)

type GameController struct {
	gameService *service.GameService
}

type IGameController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewGameController(gameService *service.GameService) *GameController {
	return &GameController{
		gameService: gameService,
	}
}

func (con *GameController) Index(c *fiber.Ctx) error {
	games, err := con.gameService.GetGames()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Game list retreived successfully!"},
		Data:     games,
	})
}

func (con *GameController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	game, err := con.gameService.GetGameByID(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The game retrieved successfully!"},
		Data:     game,
	})
}

func (con *GameController) Store(c *fiber.Ctx) error {
	req := new(request.GameRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	game, err := con.gameService.CreateGame(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The game was created successfully!"},
		Data:     game,
	})
}

func (con *GameController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(request.GameRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	game, err := con.gameService.UpdateGame(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The game was updated successfully!"},
		Data:     game,
	})
}

func (con *GameController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.gameService.DeleteGame(id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The game was deleted successfully!"},
	})
}
