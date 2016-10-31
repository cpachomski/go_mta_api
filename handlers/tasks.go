package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-echo-vue/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {


		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Instantiate new task
		var task models.Task
		//map incoming JSON body to the new task
		c.Bind(&task)
		// add new task
		id, err := models.PutTask(db, task.Name)
		//return JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		//handle errors
		} else {
			return err
		}
	}
}

func DeleteTask( db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
					"deleted": id,
				})
		//handle errors
		} else {
			return err
		}
	}
}