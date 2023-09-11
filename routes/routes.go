package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mackcoding/genji-discord-events/handlers"
)

func Init(e *echo.Echo) {

	e.GET("/", handlers.HomepageHandler)
	e.GET("/settings/", handlers.SettingsHandler)
}
