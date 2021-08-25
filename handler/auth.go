package handler

import (
	"go-prisma/service"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   24 * 60 * 60 * 30, // 1 month
		HttpOnly: true,
	}

	request := new(LoginRequest)
	c.Bind(&request)

	user, err := h.authService.ValidateUser(request.Email, request.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sess.Values["user"] = map[string]string{"email": user.Email, "name": user.Name}
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess.Values["user"] = nil
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func (h *AuthHandler) Me(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	err = LoginGuard(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, sess.Values["user"])
}
