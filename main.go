package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var server = setupSocketServer()

type LogRequest struct {
	Channel string  `json:"channel"`
	Data    logData `json:"data"`
}

func postLog(c *gin.Context) {
	var newLogRequest LogRequest

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&newLogRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	if newLogRequest.Data == nil || reflect.TypeOf(newLogRequest.Data).String() == "string" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request: Wrong log data type",
		})
		return
	}

	// Add the new album to the slice.
	addedLog := addLog(newLogRequest.Data, newLogRequest.Channel)

	server.BroadcastToRoom("/", newLogRequest.Channel, "log:create", addedLog)

	c.IndentedJSON(http.StatusCreated, addedLog)
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Authorization", "Accept", "X-Requested-With"}

	router.Use(cors.New(config))

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer server.Close()

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	router.POST("/api/logs", postLog)

	router.Run(":3004")
}
