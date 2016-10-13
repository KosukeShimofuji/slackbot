package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/kr/pretty"
	"github.com/nlopes/slack"
	"log"
	"os"
	_ "reflect"
)

// package variables
var (
	BOTNAME = "SSSBOT"
	BOTCHAN = "SSSB0T"
	TOKEN   = os.Getenv("SLACKBOT_TOKEN")
)

func info(api *slack.Client, channel, message string) {
	formated_message :=
		fmt.Sprintf("*[%s][%d][INFO]*: %s\n", BOTNAME, os.Getpid(), message)

	_, timestamp, err := api.PostMessage(channel,
		formated_message,
		slack.PostMessageParameters{Username: BOTNAME},
	)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	c := color.New(color.FgBlue)
	c.Printf("[%s][%d][%s][%s][INFO]: %s\n",
		BOTNAME, os.Getpid(), channel, timestamp, message,
	)
}

func debug(message string) {
	c := color.New(color.FgMagenta)
	c.Printf("[%s][%d][DEBUG]: %s\n", BOTNAME, os.Getpid(), message)
}

func rtmloop(api *slack.Client) {
	// initialize NewRTM
	rtm := api.NewRTM()
	go rtm.ManageConnection()
RTMLOOP:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			// fmt.Print("Event Received: ")
			// fmt.Printf("%# v\n",pretty.Formatter(msg))
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				debug(fmt.Sprintf("Message: %s", ev.Text))
			case *slack.RTMError:
				debug(fmt.Sprintf("Error: %s", ev.Error()))
			case *slack.InvalidAuthEvent:
				debug(fmt.Sprintf("Invalid credentials"))
				break RTMLOOP
			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}

func main() {
	// variable settings
	if TOKEN == "" {
		fmt.Printf("please set SSS_SLACKBOT_TOKEN environment variable\n")
		os.Exit(0)
	}

	// slackbot basic settings
	api := slack.New(TOKEN)
	logger := log.New(os.Stdout, BOTNAME+": ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	// running
	info(api, BOTCHAN, fmt.Sprintf("%sが起動しました。:grinning: ", BOTNAME))

	rtmloop(api)
	info(api, BOTCHAN, fmt.Sprintf("*FINISH %s*", BOTNAME))
}
