package routes

import (
	"net/http"

	"github.com/farhanroy/api-amalan-nahdliyin/src/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Amalan Nahdliyin!")
	})

	e.GET("/api/tahlil", controllers.FetchTahlil)

	return e
}