package repository

import (
	"context"

	"github.com/BerlitzPlatina/gf-uma/app/module/user/request"
	"github.com/BerlitzPlatina/gf-uma/internal/bootstrap/database"
	"github.com/BerlitzPlatina/gf-uma/internal/ent"
	"github.com/BerlitzPlatina/gf-uma/internal/ent/user"
)

type UserRepository struct {
	DB *database.Database
}

type IUserRepository interface {
	GetUsers() ([]*ent.User, error)
	GetUserByID(id int) (*ent.User, error)
	CreateUser(request request.UserRequest) (*ent.User, error)
	UpdateUser(id int, request request.UserRequest) (*ent.User, error)
	DeleteUser(id int) error
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{
		DB: database,
	}
}

func (s *UserRepository) GetUsers() ([]*ent.User, error) {
	return s.DB.Ent.User.Query().Order(ent.Asc(user.FieldID)).All(context.Background())
}

func (s *UserRepository) GetUserByID(id int) (*ent.User, error) {
	return s.DB.Ent.User.Query().Where(user.IDEQ(id)).First(context.Background())
}

func (s *UserRepository) CreateUser(request request.UserRequest) (*ent.User, error) {
	return s.DB.Ent.User.Create().
		SetUsername(request.Name).
		SetPassword(request.Icon).
		Save(context.Background())
}

func (s *UserRepository) UpdateUser(id int, request request.UserRequest) (*ent.User, error) {
	return s.DB.Ent.User.UpdateOneID(id).
		SetUsername(request.Name).
		SetPassword(request.Icon).
		Save(context.Background())
}

func (s *UserRepository) DeleteUser(id int) error {
	return s.DB.Ent.User.DeleteOneID(id).Exec(context.Background())
}

func (s *UserRepository) GetUserByParams(params ent.User) (*ent.User, error) {
	return s.DB.Ent.User.Query().
		Where(user.Password(params.Username)).
		Where(user.Password(params.Password)).
		First(context.Background())
}
