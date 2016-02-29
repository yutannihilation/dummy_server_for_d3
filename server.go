package main

import (
	"net/http"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
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
func hello(c *gin.Context) {
	c.JSON(http.StatusOK, dummy_data())
}

func dummy_data() Record {
	return Record{Timestamp: myTime{time.Now()}, Value: rand.Intn(100)}
}

func main() {
	// Echo instance
	r := gin.Default()

	// Routes
	r.GET("/data", hello)
	r.StaticFile("/", "index.html")

	// Start server
	r.Run(":1323")
}
