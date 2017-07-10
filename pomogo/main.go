package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0xAX/notificator"
	"github.com/Sirupsen/logrus"
)

const (
	workTitle = "WORK TIME D:"
	workMsg   = "Let your hobbies and focus on your task ;)"
	restTitle = "REST TIME :D"
	restMsg   = "Let your task and focus on your hobbies :D"
)

var (
	workTime = "0"
	restTime = "0"
)

func main() {

	if !(len(os.Args) > 1) || len(os.Args) > 3 {
		fmt.Println("[Usage]: ./pomogo <workTime> <restTime>")
		os.Exit(1)
	} else {
		workTime = os.Args[1]
		restTime = os.Args[2]
	}

	notifier := notificator.New(notificator.Options{
		DefaultIcon: "shout.png",
		AppName:     "PomoGo",
	})

	for {
		sleepAndNotify(workTime, workTitle, workMsg, notifier)
		sleepAndNotify(restTime, restTitle, restMsg, notifier)
	}
}

// sleepAndNotify abstract the notification creation process.
func sleepAndNotify(t, title, msg string, notifier *notificator.Notificator) {

	if err := notifier.Push(title, msg, "", notificator.UR_NORMAL); err != nil {
		logrus.Warn(err)
	}

	// Get
	waitTime, err := time.ParseDuration(t + "s")
	if err != nil {
		logrus.Warn(err)
	}

	time.Sleep(waitTime)
}
