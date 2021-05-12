package main

import (
	"github.com/bwmarrin/discordgo"
)

func Reply(cid string, m string, args []string) {
	s.ChannelMessageSend(cid, m)
}

func Ereply(cid string, m *discordgo.MessageEmbed) {
	s.ChannelMessageSendEmbed(cid, m)
}
