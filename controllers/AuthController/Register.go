package authcontroller

import (
	"net/http"
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputRegister struct {
	Username string `validate:"required,min=3"`
	Password string `validate:"required,min=5"`
}

func Register(c echo.Context) error {
	db := database.DBManager()
	var user models.User
	var input inputRegister

	if err := c.Bind(&input); err != nil {
		return err
	}

	err := registerValidate(c, input)
	if err != nil {
		return err
	}

	user.Username = input.Username
	user.Password = helpers.Hash(input.Password)

	result := db.Create(&user)
	if result.Error != nil {
		return c.JSON(500, result.Error)
	}

	return c.JSON(http.StatusOK, &user)
}

func registerValidate(c echo.Context, input inputRegister) error {

	v := validator.New()

	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
