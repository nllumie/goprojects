package handlers

import (
	"context"
	"net/http"
	"strconv"

	"hw3/db"
	"hw3/models"

	"github.com/labstack/echo/v4"
)

func GetAllSchedules(c echo.Context) error {
	rows, err := db.Pool.Query(
		context.Background(),
		`SELECT schedule_id, group_id, subject, time_slot FROM schedule`,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "database error",
		})
	}
	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var s models.Schedule
		rows.Scan(
			&s.ScheduleID,
			&s.GroupID,
			&s.Subject,
			&s.TimeSlot,
		)
		schedules = append(schedules, s)
	}

	return c.JSON(http.StatusOK, schedules)
}

func GetScheduleByGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid group id",
		})
	}

	rows, err := db.Pool.Query(
		context.Background(),
		`SELECT schedule_id, group_id, subject, time_slot 
		 FROM schedule 
		 WHERE group_id = $1`,
		groupID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "database error",
		})
	}
	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var s models.Schedule
		rows.Scan(
			&s.ScheduleID,
			&s.GroupID,
			&s.Subject,
			&s.TimeSlot,
		)
		schedules = append(schedules, s)
	}

	return c.JSON(http.StatusOK, schedules)
}
