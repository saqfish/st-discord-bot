package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/barthr/newsapi"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session
var c *newsapi.Client
var ars []newsapi.Article
var count int

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Args[1])
	if err != nil {
		log.Fatal("New bot error")
	}
	c = newsapi.NewClient(os.Args[2], newsapi.WithHTTPClient(http.DefaultClient))
}

func news(cid string, m string) {
	hl, _ := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"google-news", "NPR"},
	})
	for _, a := range hl.Articles {
		ars = append(ars, a)
	}
	msg := fmt.Sprintf("Got %d articles", len(ars))
	reply(cid, msg)
}

func prev(cid string, s string) {
	if count == 0 {
		reply(cid, "No more")
		return
	} else {
		count--
	}
	if len(ars) == 0 {
		reply(cid, "Get articles first fool")
		return
	}
	Ereply(cid, atoe(ars[count]))
	return
}

func next(cid string, s string) {
	if count >= len(ars) {
		reply(cid, "No more")
	} else {
		count++
	}
	if len(ars) == 0 {
		reply(cid, "Get articles first fool")
	}
	Ereply(cid, atoe(ars[count]))
}

func atoe(a newsapi.Article) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         a.URL,
		Type:        discordgo.EmbedTypeRich,
		Title:       a.Title,
		Description: a.Description,
		Author:      &discordgo.MessageEmbedAuthor{Name: a.Author},
	}
}

func repeat(cid string, s string) {
	reply(cid, s)
}

func reply(cid string, m string) {
	s.ChannelMessageSend(cid, m)
}

func Ereply(cid string, m *discordgo.MessageEmbed) {
	log.Println("sending embeded")
	fmt.Println(m)
	s.ChannelMessageSendEmbed(cid, m)
}

func main() {
	commands := map[string]func(cid string, m string){
		"taco": repeat,
		"get":  news,
		"next": next,
		"prev": prev,
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if r, ok := commands[m.Content]; ok {
			r(m.ChannelID, m.Content)
		}
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	s.Close()
}
