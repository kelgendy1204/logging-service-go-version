package main

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func setupSocketServer() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnEvent("/logs", "joinChannel", func(s socketio.Conn, channelMsg string) {
		s.Join(channelMsg)
	})

	server.OnEvent("/logs", "log:delete", func(s socketio.Conn, channelMsg string) {
		clearChannelLogs(channelMsg)
	})

	server.OnEvent("/logs", "log:list", func(s socketio.Conn, channelMsg string) []Log {
		fmt.Println("channelMsg")
		logs := getChannelLogs(channelMsg)
		return logs
	})

	return server
}
