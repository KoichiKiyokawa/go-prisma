//go:generate go run github.com/prisma/prisma-client-go generate
package main

import (
	"go-prisma/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router.SetupRoutes(e, client)

	e.Logger.Fatal(e.Start(":8080"))
}
