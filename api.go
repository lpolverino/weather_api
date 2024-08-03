package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	ListenAddr string
}

type Query struct {
	Location     string `json:"location" binding:"required"`
	StartingDate string `json:"date_1" bson:",omitempty"`
	EndingDate   string `json:"date_2" bson:",omitempty"`
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
	}
}

func (a *APIServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", func(c echo.Context) error {
		queryUrl := "https://www.example.com/https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
		newQuery := Query{}
		err := c.Bind(&newQuery)
		if err != nil {
			return c.String(http.StatusBadRequest, "Cannot process query")
		}
		queryUrl += newQuery.Location

		fmt.Printf("%+v\n", newQuery)
		fmt.Printf("%d", len(newQuery.EndingDate))

		if len(newQuery.StartingDate) != 0 && len(newQuery.EndingDate) != 0 {
			queryUrl += "/" + newQuery.StartingDate + "/" + newQuery.EndingDate
		} else if !(len(newQuery.StartingDate) == 0 && len(newQuery.EndingDate) == 0) {
			return c.String(http.StatusBadRequest, "cannot process query")
		}

		return c.String(http.StatusAccepted, queryUrl)
	})

	e.Logger.Fatal(e.Start(":" + a.ListenAddr))
}
