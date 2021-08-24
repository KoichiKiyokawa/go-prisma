package router

import (
	"go-prisma/prisma/db"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, dbClient *db.PrismaClient) {
	UserRouter(e, dbClient)
}
