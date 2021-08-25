package router

import (
	"go-prisma/handler"
	"go-prisma/prisma/db"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, dbClient *db.PrismaClient) {
	userHandler := handler.NewUserHandler(dbClient)
	g := e.Group("/users/")
	g.GET("", userHandler.Index)
	g.GET(":id", userHandler.Show)
	g.POST("", userHandler.Create)
	g.PATCH(":id", userHandler.Update)
	g.DELETE(":id", userHandler.Delete)
}
