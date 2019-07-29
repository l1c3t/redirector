package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//HWorld just another path to helloWorld
func HWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
