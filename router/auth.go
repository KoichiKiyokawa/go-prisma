package router

import (
	"go-prisma/handler"
	"go-prisma/prisma/db"
	"go-prisma/service"

	"github.com/labstack/echo/v4"
)

func AuthRouter(e *echo.Echo, client db.PrismaClient) {
	authService := service.NewAuthService(client)
	authHandler := handler.NewAuthHandler(authService)

	g := e.Group("/auth/")
	g.POST("login", authHandler.Login)
	g.POST("logout", authHandler.Logout)
	g.GET("me", authHandler.Me)
}
