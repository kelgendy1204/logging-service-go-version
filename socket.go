package main

import (
	socketio "github.com/googollee/go-socket.io"
)

func setupSocketServer() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnEvent("/", "joinChannel", func(s socketio.Conn, channelMsg string) {
		s.Join(channelMsg)
	})

	server.OnEvent("/", "log:delete", func(s socketio.Conn, channelMsg string) {
		clearChannelLogs(channelMsg)
	})

	server.OnEvent("/", "log:list", func(s socketio.Conn, channelMsg string) []Log {
		logs := getChannelLogs(channelMsg)
		return logs
	})

	return server
}
