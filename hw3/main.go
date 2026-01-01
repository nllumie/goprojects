package main

import (
	"hw3/db"
	"hw3/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Connect()

	e := echo.New()

	e.GET("/student/:id", handlers.GetStudentByID)
	e.GET("/all_class_schedule", handlers.GetAllSchedules)
	e.GET("/schedule/group/:id", handlers.GetScheduleByGroup)

	e.Logger.Fatal(e.Start(":8080"))
}
