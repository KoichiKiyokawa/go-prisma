//go:generate go run github.com/prisma/prisma-client-go generate
package main

import (
	"go-prisma/prisma/db"
	"go-prisma/router"
	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secrethogehogehogehoge"))))

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

	log.Print(e.Routes())
	e.Logger.Fatal(e.Start(":8080"))
}
