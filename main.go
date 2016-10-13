package main

import (
	"fmt"
	_ "github.com/fatih/color"
	_ "github.com/kr/pretty"
	"github.com/nlopes/slack"
	"log"
	"os"
	_ "reflect"
)

// package variables
var (
	BOTNAME = "3sb0t"
	BOTCHAN = "3sb0tr00m"
	TOKEN   = os.Getenv("SSS_SLACKBOT_TOKEN")
)

func speak(api *slack.Client, message string) {
	// send startup message
	msg := fmt.Sprintf("%s[%d]: %s\n", BOTNAME, os.Getpid(), message)
	channelID, timestamp, err := api.PostMessage(BOTCHAN, msg, slack.PostMessageParameters{Username: BOTNAME})

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}

func rtmloop(api *slack.Client) {
	// initialize NewRTM
	rtm := api.NewRTM()
	go rtm.ManageConnection()
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			// fmt.Print("Event Received: ")
			// fmt.Printf("%# v\n",pretty.Formatter(msg))
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Printf("Message: %s\n", ev.Text)
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
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

	speak(api, fmt.Sprintf("start up %s", BOTNAME))

	rtmloop(api)

	speak(api, fmt.Sprintf("end %s", BOTNAME))
}

