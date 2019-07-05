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
	Number string  `faker:"cc_number" json:"number"`
	Period string  `faker:"month_name" json:"period"`
	Seller string  `faker:"name" json:"seller,omitempty"`
	Set    float64 `faker:"amount" json:"set"`
	Buyer  string  `faker:"name" json:"buyer,omitempty"`
}

// LotteryResponse struct
type LotteryResponse struct {
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

	e.GET("/lottery", func(c echo.Context) error {
		ll := []Lottery{}
		for i := 0; i < 10; i++ {
			l := Lottery{}
			_ = faker.FakeData(&l)
			ll = append(ll, l)
		}

		return c.JSON(http.StatusOK, LotteryResponse{
			Data: ll,
		})
	})

	e.GET("/inventory", func(c echo.Context) error {
		ll := []Lottery{}
		for i := 0; i < 10; i++ {
			l := Lottery{}
			_ = faker.FakeData(&l)
			ll = append(ll, l)
		}

		return c.JSON(http.StatusOK, LotteryResponse{
			Data: ll,
		})
	})

	e.GET("/lottery/search", func(c echo.Context) error {
		l := Lottery{}
		err := faker.FakeData(&l)
		if err != nil {
			log.Fatal("Unable to create fake data")
		}

		return c.JSON(http.StatusOK, LotteryResponse{
			Data: []Lottery{l},
		})
	})

	e.GET("/orders", func(c echo.Context) error {
		ll := []Lottery{}
		for i := 0; i < 10; i++ {
			l := Lottery{}
			_ = faker.FakeData(&l)
			ll = append(ll, l)
		}

		return c.JSON(http.StatusOK, LotteryResponse{
			Data: ll,
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

		return c.JSON(http.StatusOK, LotteryResponse{
			Data: ll,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
