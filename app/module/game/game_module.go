package game

import (
	"github.com/BerlitzPlatina/gf-uma/app/module/game/controller"
	"github.com/BerlitzPlatina/gf-uma/app/module/game/repository"
	"github.com/BerlitzPlatina/gf-uma/app/module/game/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type GameRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// Register bulkly
var NewGameModule = fx.Options(
	// Register Repository & Service
	fx.Provide(repository.NewGameRepository),
	fx.Provide(service.NewGameService),

	// Regiser Controller
	fx.Provide(controller.NewController),

	// Register Router
	fx.Provide(NewGameRouter),
)

// Router methods
func NewGameRouter(fiber *fiber.App, controller *controller.Controller) *GameRouter {
	return &GameRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *GameRouter) RegisterGameRoutes() {
	// Define controllers
	gameController := r.Controller.Game

	// Define routes
	r.App.Route("/games", func(router fiber.Router) {
		router.Get("/", gameController.Index)
		router.Get("/game/:id", gameController.Show)
		router.Post("/game/", gameController.Store)
		router.Patch("/game/:id", gameController.Update)
		router.Delete("/game/:id", gameController.Destroy)
	})
}
