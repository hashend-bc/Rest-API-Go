package main

import (
	"go-rest-api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/items", handlers.AddItem)
	e.GET("/items", handlers.GetItems)
	e.PUT("/items/:id", handlers.UpdateItem)
	e.DELETE("/items/:id", handlers.DeleteItem)

	e.Logger.Fatal(e.Start(":8080"))
}