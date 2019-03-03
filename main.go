package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func postUsers(c echo.Context) error {
  return nil
}

func postAuthenticate(c echo.Context) error {
  return nil
}

func getChannels(c echo.Context) error {
  return nil
}

func getMessages(c echo.Context) error {
  return nil
}

func postAddChannel(c echo.Context) error {
  return nil
}

func postAddMessage(c echo.Context) error {
  return nil
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
  e.GET("/api/channels", getChannels)
  e.GET("/api/messages", getMessages)
  e.POST("/api/users", postUsers)
  e.POST("/api/authenticate", postAuthenticate)
  e.POST("/api/channels", postAddChannel)
  e.POST("/api/messages", postAddMessage)
	e.Logger.Fatal(e.Start(":8000"))
}
