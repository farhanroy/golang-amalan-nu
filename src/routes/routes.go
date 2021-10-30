package routes

import (

	"github.com/farhanroy/api-amalan-nahdliyin/src/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("landing/index.html")
	})

	e.GET("/docs", func(c echo.Context) error {
		return c.File("docs/index.html")
	})

	e.GET("/api/tahlil", controllers.FetchTahlil)

	e.GET("/api/shalawat", controllers.FetchShalawat)

	e.GET("/api/haul", controllers.FetchHaul)

	e.GET("/api/istighotsah", controllers.FetchIstighotsah)

	return e
}