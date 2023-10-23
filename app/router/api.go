package router

import (
	"github.com/BerlitzPlatina/gf-uma/app/module/article"
	"github.com/BerlitzPlatina/gf-uma/app/module/game"
	"github.com/BerlitzPlatina/gf-uma/app/module/user"
	"github.com/BerlitzPlatina/gf-uma/storage"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App           fiber.Router
	ArticleRouter *article.ArticleRouter
	GameRouter    *game.GameRouter
	UserRouter    *user.UserRouter
}

func NewRouter(fiber *fiber.App, articleRouter *article.ArticleRouter, gameRouter *game.GameRouter, userRouter *user.UserRouter) *Router {
	return &Router{
		App:           fiber,
		ArticleRouter: articleRouter,
		GameRouter:    gameRouter,
		UserRouter:    userRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	r.App.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(string(example))
	})

	// Register routes of modules
	r.ArticleRouter.RegisterArticleRoutes()
	r.GameRouter.RegisterGameRoutes()
	r.UserRouter.RegisterUserRoutes()
}
