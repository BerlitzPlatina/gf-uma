package game

import (
	"github.com/efectn/fiber-boilerplate/app/module/game/controller"
	"github.com/efectn/fiber-boilerplate/app/module/game/repository"
	"github.com/efectn/fiber-boilerplate/app/module/game/service"
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
		router.Get("/:id", gameController.Show)
		router.Post("/", gameController.Store)
		router.Patch("/:id", gameController.Update)
		router.Delete("/:id", gameController.Destroy)
	})
}
