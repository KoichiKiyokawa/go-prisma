package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoginGuard(c echo.Context) *echo.HTTPError {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if sess.Values["user"] == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return nil
}
