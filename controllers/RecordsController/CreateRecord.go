package recordscontroller

import (
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputCreateRecord struct {
	Record string `json:"record" validate:"required"`
}

func CreateRecord(c echo.Context) error {
	var record models.Record
	var input inputCreateRecord
	db := database.DBManager()

	c.Bind(&input)
	err := recordValidate(c, input)
	if err != nil {
		return err
	}
	record.UserId = helpers.User(c).ID
	record.Record = input.Record

	db.Create(&record)
	return c.JSON(http.StatusOK, map[string]bool{
		"isCreated": true,
	})
}
func recordValidate(c echo.Context, input inputCreateRecord) error {

	v := validator.New()

	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
