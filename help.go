package main

import "github.com/nlopes/slack"

var (
	commands = map[string]string{
		"help":   "Displays all of the help commands.",
		"dakoku": "I will punch a clock.",
	}
)

func (b *Bot) help() []slack.Attachment {
	fields := make([]slack.AttachmentField, 0)

	for k, v := range commands {
		fields = append(fields, slack.AttachmentField{
			Title: "@" + botName + " " + k,
			Value: v,
		})
	}

	attachment := []slack.Attachment{slack.Attachment{
		Pretext: botName + " Command List",
		Color:   "#B733FF",
		Fields:  fields,
	}}
	return attachment
}
