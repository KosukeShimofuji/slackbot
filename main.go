package main

import (
	"fmt"
	"log"
	"os"

        //"github.com/fatih/color"
	"github.com/nlopes/slack"
	//debug
	//"reflect"
	"github.com/kr/pretty"
)

var (
	BOTNAME = "sss-bot"
	BOTCHAN = "sss-bot-chan"
	TOKEN   = os.Getenv("SSS_SLACKBOT_TOKEN")
)

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

	// send startup message
	msg := fmt.Sprintf("START UP %s [%d]\n", BOTNAME, os.Getpid())
	channelID, timestamp, err := api.PostMessage(BOTCHAN, msg, slack.PostMessageParameters{})

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

        // initialize NewRTM
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
                        fmt.Printf("%# v\n",pretty.Formatter(msg))
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
                            // Ignore hello
			case *slack.ConnectedEvent:
				fmt.Println("Infos:", ev.Info)
				fmt.Println("Connection counter:", ev.ConnectionCount)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", BOTCHAN))
			case *slack.Message:
				fmt.Printf("Message: %v\n", ev)
			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)
			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)
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



