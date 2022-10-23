package recordscontroller

import (
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRecord(c echo.Context) error {
	var record models.Record
	db := database.DBManager()
	db.Where("`user_id` = " + strconv.Itoa(helpers.User(c).ID)).Order("created_at desc").First(&record)
	return c.JSON(http.StatusOK, map[string]models.Record{
		"record": record,
	})
}
