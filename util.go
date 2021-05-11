package main

import (
	"github.com/barthr/newsapi"
	"github.com/bwmarrin/discordgo"
)

func AtoE(a newsapi.Article) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         a.URL,
		Type:        discordgo.EmbedTypeRich,
		Title:       a.Title,
		Description: a.Description,
		Author:      &discordgo.MessageEmbedAuthor{Name: a.Author},
	}
}
