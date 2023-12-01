package handler

import (
	"log"
	"net/http"
	"sqlite/pkg/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID          string `gorm:"primaryKey"`
		ConnectedID string
	}

	GetUsersRequest  struct{}
	GetUsersResponse struct {
		Users []Task `json:"users"`
	}

	CreateUserRequest struct {
		ID string `json:"id"`
	}
	CreateUserResponse struct{}

	UpdateUserRequest struct {
		ID          string `param:"id"`
		ConnectedID string `json:"connected_id"`
	}
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
		ConnectedID: "",
	}

	if err := h.db.Create(dbUser).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to create user")
	}

	return c.JSON(http.StatusOK, dbUser)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	user := &UpdateUserRequest{}
	if err := c.Bind(user); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Print("UserID:" + user.ID)
	log.Print("ConnectedID:" + user.ConnectedID)

	upUser := &database.User{
		ID:          user.ID,
		ConnectedID: user.ConnectedID,
	}

	if err := h.db.Where("id = ?", user.ID).Save(&upUser).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to update user")
	}

	log.Print("update user")
	return c.JSON(http.StatusOK, "success")
}
