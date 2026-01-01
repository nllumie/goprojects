package handlers

import (
	"context"
	"net/http"
	"strconv"

	"hw3/db"
	"hw3/models"

	"github.com/labstack/echo/v4"
)

func GetStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid student id",
		})
	}

	var student models.Student

	query := `
		SELECT 
			s.id,
			s.firstname,
			s.email,
			s.mobile,
			s.gender,
			s.group_id,
			g.group_name
		FROM students s
		LEFT JOIN groups g ON s.group_id = g.group_id
		WHERE s.id = $1
	`

	err = db.Pool.QueryRow(context.Background(), query, id).
		Scan(
			&student.ID,
			&student.FirstName,
			&student.Email,
			&student.Mobile,
			&student.Gender,
			&student.GroupID,
			&student.GroupName,
		)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "student not found",
		})
	}

	return c.JSON(http.StatusOK, student)
}
