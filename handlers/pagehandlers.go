package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HomepageHandler(c echo.Context) error {
	fmt.Printf("HomepageHandler")
	return c.File("./static/htmx/index.htmx")
}

// settings.xhtm
func SettingsHandler(c echo.Context) error {
	return c.File("./static/htmx/settings/index.htmx")
}
