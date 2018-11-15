package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

var (
	botID   string
	botName string
)

func bot() {
	token := os.Getenv("SLACK_API_TOKEN")
	b := newBot(token)

	go b.rtm.ManageConnection()

	done := make(chan struct{})
	go func() {
		defer close(done)

		for msg := range b.rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				botName = ev.Info.User.Name

			case *slack.MessageEvent:
				user := ev.User
				text := ev.Text
				channel := ev.Channel
				if ev.Type == "message" && strings.HasPrefix(text, "<@"+botID+">") {
					b.handleResponse(user, text, channel)
				}
			case *slack.DisconnectedEvent:
				return
			}
		}
	}()
	<-done
	return
}

func newBot(token string) *Bot {
	bot := new(Bot)
	bot.api = slack.New(token)
	bot.api.SetDebug(true)
	bot.rtm = bot.api.NewRTM()
	return bot
}

func (b *Bot) handleResponse(user, text, channel string) {
	commandArray := strings.Fields(text)
	cmd := "help"

	if len(commandArray) > 0 {
		cmd = commandArray[1]
	}

	var attachments []slack.Attachment
	var err error
	switch cmd {
	case "dakoku":
		attachments, err = b.punchClock()
	case "help":
		attachments = b.help()
	default:
		attachments = b.help()
	}

	if err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error... %s", cmd, err), channel))
		return
	}

	params := slack.PostMessageParameters{
		Attachments: attachments,
		Username:    botName,
	}

	_, _, err = b.api.PostMessage(channel, "", params)
	if err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error... %s", cmd, err), channel))
		return
	}
}
