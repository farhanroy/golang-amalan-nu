package controllers

import (
	"net/http"

	"github.com/farhanroy/api-amalan-nahdliyin/src/models"
	"github.com/labstack/echo"
)

func FetchShalawat(c echo.Context) error {
	result, err := models.FetchShalawat()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
