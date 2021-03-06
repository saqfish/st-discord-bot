package main

import (
	"fmt"

	"github.com/barthr/newsapi"
	"github.com/bwmarrin/discordgo"
)

func ArticleToEmbed(a newsapi.Article) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         a.URL,
		Type:        discordgo.EmbedTypeRich,
		Title:       a.Title,
		Description: a.Description,
		Author:      &discordgo.MessageEmbedAuthor{Name: a.Author},
	}
}

func ComicToEembed(c Comic) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         c.Img,
		Type:        discordgo.EmbedTypeRich,
		Title:       c.Safe_title,
		Image:       &discordgo.MessageEmbedImage{URL: c.Img},
		Description: c.Alt,
	}
}

func JokeToEmbed(j Joke) *discordgo.MessageEmbed {
	p := fmt.Sprintf("||%s||", j.Punchline)
	return &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       j.Setup,
		Description: p,
	}
}
