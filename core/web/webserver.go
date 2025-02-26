package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func run() error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
		})
	})

	// Start the Gin server in a separate goroutine to prevent blocking
	go func() {
		if err := r.Run(); err != nil {
			panic(err) // Handle the error if the server fails to start
		}
	}()
	return nil
}

// Start used to init the db into mem as early as possible
func Start() {
	log.Println("starting webserver")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
