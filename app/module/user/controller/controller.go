package controller

import "github.com/BerlitzPlatina/gf-uma/app/module/user/service"

type Controller struct {
	User *UserController
	Auth *AuthController
}

func NewController(userService *service.UserService) *Controller {
	return &Controller{
		User: &UserController{userService: userService},
		Auth: &AuthController{userService: userService},
	}
}
