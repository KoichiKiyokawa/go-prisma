package handler

import (
	"context"
	"go-prisma/prisma/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	client *db.PrismaClient
}

// TODO: この構造体を使えないか。。。
type UserResponse struct {
	Password string `json:"-"`
	db.UserModel
}

func NewUserHandler(client *db.PrismaClient) *UserHandler {
	return &UserHandler{client}
}

func (h *UserHandler) Index(c echo.Context) error {
	users, err := h.client.User.FindMany().Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "error get users")
	}

	// TODO: omit password
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Show(c echo.Context) error {
	user, err := h.client.User.FindUnique(db.User.ID.Equals(c.Param("id"))).Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c echo.Context) error {
	var user db.UserModel
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	created, err := h.client.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(string(hashedPassword)),
	).Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, created)
}

func (h *UserHandler) Update(c echo.Context) error {
	var user db.UserModel
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	updatedUser, err := h.client.User.UpsertOne(
		db.User.ID.Equals(c.Param("id")),
	).Exec(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) Delete(c echo.Context) error {
	// h.client.User
	return c.JSON(http.StatusOK, map[string]string{})
}
