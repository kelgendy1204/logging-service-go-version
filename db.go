package main

import "time"
import "strconv"

type logData interface {}

type Log struct {
    Time string     `json:"time"`
    Data logData    `json:"data"`
}

const MaxLogs = 100;
var logs = make(map[string][]Log)

func addLog(data logData, channel string) Log {
	if logs[channel] == nil {
		logs[channel] = make([]Log, 0)
	}

    currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
    addedLog := Log { Time: currentTimestamp, Data: data }

    logs[channel] = append(logs[channel], addedLog)

    // Check if the logs slice exceeds its maximum capacity
    if len(logs[channel]) > MaxLogs {
        // Remove the oldest log entry
        logs[channel] = logs[channel][1:]
    }

    return addedLog
}

func getChannelLogs(channel string) []Log {
    currentLog := logs[channel]

	if currentLog != nil {
		return currentLog
	}

	return []Log{}
}

func clearChannelLogs(channel string) {
	delete(logs, channel)
}
