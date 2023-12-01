package handler

import (
	"net/http"
	"sqlite/pkg/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID          string `gorm:"primaryKey"`
		ConnectedID *string
	}

	GetUsersRequest  struct{}
	GetUsersResponse struct {
		Users []Task `json:"users"`
	}

	CreateUserRequest struct {
		ID string `json:"id"`
	}
	CreateUserResponse struct{}
)

type userHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{db: db}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	users := []User{}

	if err := h.db.Table("users").Find(&users).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to get tasks")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &CreateUserRequest{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	dbUser := &database.User{
		ID:          user.ID,
		ConnectedID: nil,
	}

	if err := h.db.Create(dbUser).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to create user")
	}

	return c.JSON(http.StatusOK, dbUser)
}
