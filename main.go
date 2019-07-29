package main

import (
	"redir/handlers"

	"github.com/labstack/echo"
)

func main() {

	/* пример запросов:
	1. string	http://127.0.0.1:8080/redir/string?shrt=https://302.su/r/8y
	2. json	http://127.0.0.1:8080/redir/json?shrt=https://302.su/r/8y
	*/
	e := echo.New()
	e.GET("/tst", handlers.HWorld)
	e.GET("/redir/:data", handlers.Redir)
	e.Logger.Fatal(e.Start(":80"))
}
