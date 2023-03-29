package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

/* const io = req.app.get('socketio');

   const addedLog = addLog(log, channel);
   io.to(channel).emit('log:create', addedLog);

   return res.status(200).end(); */

type LogRequest struct {
    Channel string  `json:"channel"`
    Data logData `json:"data"`
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

    c.IndentedJSON(http.StatusCreated, addedLog)
}

func main() {
    router := gin.Default()

    router.POST("/api/logs", postLog)

    router.Run("localhost:8070")
}
