package recordscontroller

import (
	"net/http"
	"setup/database"
	"setup/models"

	"github.com/labstack/echo/v4"
)

func GetRecord(c echo.Context) error {
	var record models.Record
	db := database.DBManager()
	db.Order("created_at desc").First(&record)
	return c.JSON(http.StatusOK, map[string]models.Record{
		"record": record,
	})
}
