package main

import (
	"log"
	"net/http"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Lottery struct
type Lottery struct {
	Number string `faker:"cc_number" json:"number"`
	Period string `faker:"month_name" json:"period"`
	Seller string `faker:"name" json:"seller,omitempty"`
}

// History struct
type History struct {
	Data []Lottery `json:"data"`
}

// SearchResponse struct
type SearchResponse struct {
	Data []Lottery `json:"data"`
}

// NewResponse struct
type NewResponse struct {
	Acknowledge bool `json:"acknowledge"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.GET("/lottery/search", func(c echo.Context) error {
		l := Lottery{}
		err := faker.FakeData(&l)
		if err != nil {
			log.Fatal("Unable to create fake data")
		}

		return c.JSON(http.StatusOK, SearchResponse{
			Data: []Lottery{l},
		})
	})

	e.POST("/lottery/new", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewResponse{
			Acknowledge: true,
		})
	})

	e.POST("/register", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewResponse{
			Acknowledge: true,
		})
	})

	e.GET("/history", func(c echo.Context) error {
		ll := []Lottery{}
		for i := 0; i < 10; i++ {
			l := Lottery{}
			_ = faker.FakeData(&l)
			ll = append(ll, l)
		}

		return c.JSON(http.StatusOK, History{
			Data: ll,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
