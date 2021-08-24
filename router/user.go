package router

import (
	"go-prisma/handler"
	"go-prisma/prisma/db"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, dbClient *db.PrismaClient) {
	userHandler := handler.NewUserHandler(dbClient)
	g := e.Group("/users")
	g.GET("", userHandler.Index)
	g.POST("", userHandler.Create)
}
