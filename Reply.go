package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Reply(cid string, m string) {
	s.ChannelMessageSend(cid, m)
}

func Ereply(cid string, m *discordgo.MessageEmbed) {
	log.Println("sending embeded")
	fmt.Println(m)
	s.ChannelMessageSendEmbed(cid, m)
}
