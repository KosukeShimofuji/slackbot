package main

import (
    "fmt"
    "log"
    "os"

    "github.com/nlopes/slack"

    //debug
    //"reflect"
    //"github.com/kr/pretty"
)

func main() {
    api := slack.New("XXX")
    logger := log.New(os.Stdout, "sss-slackbot: ", log.Lshortfile|log.LstdFlags)
    slack.SetLogger(logger)
    //api.SetDebug(true)

    channelID, timestamp, err := api.PostMessage("sss-slackbot-channel", "Hello World", slack.PostMessageParameters{})

    if err != nil {
	fmt.Printf("%s\n", err)
	return
    }

    fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

