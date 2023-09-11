package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mackcoding/genji-discord-events/routes"
)

func main() {
	// Initialize and setup Echo's middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Customization and tweaks
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	})) // Enable gzip compression
	e.Pre(middleware.AddTrailingSlash()) // Add trailing slash

	// Enable custom error handling
	e.HTTPErrorHandler = customErrorHandler

	// Initialize the routes
	routes.Init(e)

	// Serve CSS and JS files without /static/ in the URL
	e.Static("/css", "static/css")
	e.Static("/js", "static/js")

	// Start server
	e.Logger.Fatal(e.Start("localhost:8080"))
}

func customErrorHandler(err error, c echo.Context) {
	// Check if the error is a 404 error
	if httpErr, ok := err.(*echo.HTTPError); ok && httpErr.Code == http.StatusNotFound {
		// Serve the default page here (e.g., a custom 404 HTML page)
		filePath := "./static/htmx/error.htmx"

		fmt.Printf("customErrorHandler: %s", filePath)

		// Check if the file exists
		_, err := os.Stat(filePath)
		if err == nil {
			// Serve the requested file
			c.File(filePath)
			return
		}
	}

	// For other errors or if the default page doesn't exist, you can handle them as needed
	c.String(http.StatusInternalServerError, "Internal Server Error")
}
