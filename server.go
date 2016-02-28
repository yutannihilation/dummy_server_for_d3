package main

import (
	"net/http"
	"time"

	"math/rand"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type (
	myTime struct {
		time.Time
	}
	Record struct {
		Timestamp myTime `json:"timestamp"`
		Value     int    `json:"value"`
	}
)

func (t myTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-02T15:04:05.999Z0700") + `"`), nil
}

// Handler
func hello(c *echo.Context) error {
	return c.JSON(http.StatusOK, dummy_data())
}

func dummy_data() Record {
	return Record{Timestamp: myTime{time.Now()}, Value: rand.Intn(100)}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/data", hello)
	e.Index("./index.html")

	// Start server
	e.Run(":1323")
}
