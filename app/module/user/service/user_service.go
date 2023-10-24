package service

import (
	"github.com/BerlitzPlatina/gf-uma/app/module/user/repository"
	"github.com/BerlitzPlatina/gf-uma/app/module/user/request"
	"github.com/BerlitzPlatina/gf-uma/internal/ent"
)

type UserService struct {
	Repo *repository.UserRepository
}

type IUserService interface {
	GetUsers() ([]*ent.User, error)
	GetUserByID(id int) (*ent.User, error)
	CreateUser(request request.UserRequest) (*ent.User, error)
	UpdateUser(id int, request request.UserRequest) (*ent.User, error)
	DeleteUser(id int) error
	Authenticate(id int) error
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUsers() ([]*ent.User, error) {
	return s.Repo.GetUsers()
}

func (s *UserService) GetUserByID(id int) (*ent.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(request request.UserRequest) (*ent.User, error) {
	return s.Repo.CreateUser(request)
}

func (s *UserService) UpdateUser(id int, request request.UserRequest) (*ent.User, error) {
	return s.Repo.UpdateUser(id, request)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

// auth
func (s *UserService) Authenticate(request request.LoginRequest) (*ent.User, error) {
	var params ent.User
	params.Password = request.Password
	params.Username = request.Username

	return s.Repo.GetUserByParams(params)
}
