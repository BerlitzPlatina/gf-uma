package user

import (
	"github.com/BerlitzPlatina/gf-uma/app/module/user/controller"
	"github.com/BerlitzPlatina/gf-uma/app/module/user/repository"
	"github.com/BerlitzPlatina/gf-uma/app/module/user/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type UserRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// Register bulkly
var NewUserModule = fx.Options(
	// Register Repository & Service
	fx.Provide(repository.NewUserRepository),
	fx.Provide(service.NewUserService),

	// Regiser Controller
	fx.Provide(controller.NewController),

	// Register Router
	fx.Provide(NewUserRouter),
)

// Router methods
func NewUserRouter(fiber *fiber.App, controller *controller.Controller) *UserRouter {
	return &UserRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *UserRouter) RegisterUserRoutes() {
	// Define controllers
	userController := r.Controller.User

	// Define routes
	r.App.Route("/users", func(router fiber.Router) {
		router.Get("/", userController.Index)
		router.Get("/user/:id", userController.Show)
		router.Post("/user/", userController.Store)
		router.Patch("/user/:id", userController.Update)
		router.Delete("/user/:id", userController.Destroy)
	})
}
