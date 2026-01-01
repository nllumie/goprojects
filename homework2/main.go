package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helloWorldResponse{
			Message: "Hello, World!",
		})
	})

	e.GET("\request", func(c echo.Context) error {
		name := c.QueryParam("name")

		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "parameter 'name' is empty",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "received request",
			"name":    name,
		})

	})

	e.Logger.Fatal(e.Start(":8080"))
}
