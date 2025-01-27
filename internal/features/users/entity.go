package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserID      int
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	Avatar      string
	AccessLevel string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	JWT         string
}

type UHandler interface {
	Login(echo.Context) error
	Register(echo.Context) error
	GetAllUsers(echo.Context) error
	GetUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
	ChangePassword(echo.Context) error
}

type UService interface {
	Login(string, string) (int, string, string)
	Register(User) (int, string)
	GetAllUsers(int) (int, string, []User, int)
	GetUser(int) (int, string, User)
	UpdateUser(echo.Context, int, User) (int, string)
	DeleteUser(int) (int, string)
	ChangePassword(int, User) (int, string)
}

type UQuery interface {
	Login(string) (User, error)
	Register(User) error
	IsEmailRegistered(string) bool
	GetAllUsers(int, int, int) ([]User, int, error)
	GetUser(int) (User, error)
	UpdateUser(int, User) error
	DeleteUser(int) error
	ChangePassword(int, User) error
}

type LoginValidate struct {
	Email    string `validate:"required"`
	Password string `validate:"required,min=6"`
}

type RegisterValidate struct {
	Username    string `validate:"required"`
	Email       string `validate:"required"`
	Password    string `validate:"required,min=6"`
	PhoneNumber string `validate:"required"`
}

type UpdateUserValidate struct {
	Username    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Avatar      string
	AccessLevel string
}

type ChangePasswordValidate struct {
	Password string `validate:"required,min=6"`
}
