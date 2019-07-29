package handlers

import (
	"fmt"
	"net/http"
	redirector "redirector/utils"

	"github.com/labstack/echo"
)

//Redir looking for redirects
func Redir(c echo.Context) error {
	shrtURL := c.QueryParam("shrt")
	dataType := c.Param("data")

	var redirs string

	for reURL, k := range redirector.Redirector(shrtURL) { //https://302.su/r/8y
		redirs = redirs + reURL + "\n"
		fmt.Printf("%d\t%s\n", k, reURL)
	}

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("you checked:\n%s,\nredirs: \n%s", shrtURL, redirs))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, redirector.Redirector(shrtURL))

	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "something wrong",
	})
}
