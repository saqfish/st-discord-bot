package main

import (
	"github.com/bwmarrin/discordgo"
)

func Reply(cid string, m string) {
	session.ChannelMessageSend(cid, m)
}

func Ereply(cid string, m *discordgo.MessageEmbed) {
	session.ChannelMessageSendEmbed(cid, m)
}
