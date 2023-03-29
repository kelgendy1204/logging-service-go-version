package main

import (
	"fmt"
	"testing"
)

func TestAddLogAndGetLogs(t *testing.T) {
    addLog("data", "channel");
    logs := getChannelLogs("channel");

    if len(logs) != 1 {
        t.Fatalf("Log is not added");
    }
}

func TestLogsLimitPerChannel(t *testing.T) {
    for i := 0; i < 102; i++ {
        addLog(fmt.Sprintf("data %d", i), "channel");
    }

    logs := getChannelLogs("channel");

    if len(logs) > 100 {
        t.Fatalf("Logs exceeded the limit");
    }

    if logs[0].data != "data 2" {
        t.Fatalf("When logs exceeded limit the data isn't replaced correctly");
    }
}

func TestClearLogsForChannel(t *testing.T) {
    for i := 0; i < 102; i++ {
        addLog(fmt.Sprintf("data %d", i), "channel");
    }

    logs1 := getChannelLogs("channel");

    logs2 := getChannelLogs("channel2");
    clearChannelLogs("channel2")
    logs2 = getChannelLogs("channel2");

    if len(logs1) == 0 {
        t.Fatalf("Clear the wrong logs channel");
    }

    if len(logs2) != 0 {
        t.Fatalf("Not cleared the channel correctly");
    }
}
