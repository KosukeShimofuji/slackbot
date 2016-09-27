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

    channels, err := api.GetChannels(false)
    if err != nil {
	fmt.Printf("%s\n", err)
	return
    }

    for _, channel := range channels {
        fmt.Println(channel.ID, channel.Name)

//    slack.Channel{
//        groupConversation: slack.groupConversation{
//            conversation: slack.conversation{
//                ID:                 "C2G9S2KPY",
//                Created:            1474945736,
//                IsOpen:             false,
//                LastRead:           "",
//                Latest:             (*slack.Message)(nil),
//                UnreadCount:        0,
//                UnreadCountDisplay: 0,
//            },
//            Name:       "sss-slackbot-channel",
//            Creator:    "U03RAJ6BK",
//            IsArchived: false,
//            Members:    {"U03RAJ6BK"},
//            NumMembers: 1,
//            Topic:      slack.Topic{},
//            Purpose:    slack.Purpose{},
//        },
//        IsChannel: true,
//        IsGeneral: false,
//        IsMember:  false,
//    }

        //fmt.Printf("%# v\n", pretty.Formatter(channel))
        //fmt.Println(reflect.TypeOf(channel))
	//fmt.Println(channel)
    }

    //rtm := api.NewRTM()
    //go rtm.ManageConnection()
    //rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "sss-slackbot-channel"))
}


