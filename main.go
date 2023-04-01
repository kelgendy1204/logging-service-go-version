package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var server = setupSocketServer()

type LogRequest struct {
	Channel string  `json:"channel"`
	Data    logData `json:"data"`
}

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
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

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer server.Close()

    router.Use(GinMiddleware("http://localhost:3000"))

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	router.POST("/api/logs", postLog)

	router.Run(":3004")
}
