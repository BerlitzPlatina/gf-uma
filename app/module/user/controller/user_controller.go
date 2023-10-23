package controller

import (
	"strconv"

	"github.com/BerlitzPlatina/gf-uma/app/module/user/request"
	"github.com/BerlitzPlatina/gf-uma/app/module/user/service"
	"github.com/BerlitzPlatina/gf-uma/utils/response"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
}

type IUserController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (con *UserController) Index(c *fiber.Ctx) error {
	users, err := con.userService.GetUsers()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User list retreived successfully!"},
		Data:     users,
	})
}

func (con *UserController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user, err := con.userService.GetUserByID(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The user retrieved successfully!"},
		Data:     user,
	})
}

func (con *UserController) Store(c *fiber.Ctx) error {
	req := new(request.UserRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	user, err := con.userService.CreateUser(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The user was created successfully!"},
		Data:     user,
	})
}

func (con *UserController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(request.UserRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	user, err := con.userService.UpdateUser(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The user was updated successfully!"},
		Data:     user,
	})
}

func (con *UserController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.userService.DeleteUser(id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The user was deleted successfully!"},
	})
}
