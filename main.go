package main

import (
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

func main() {
	commands := map[string]func(cid string, m string){
		"taco": Reply,
		"get":  Get,
		"next": Next,
		"prev": Prev,
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
