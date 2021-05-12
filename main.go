package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
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
	if len(os.Args) < 3 {
		log.Fatal("go run main.go [DISCORD KEY] [NEWS-API KEY]")
		os.Exit(1)
	}
	if err != nil || len(os.Args) < 3 {
		log.Fatal("New bot error")
		os.Exit(1)
	}
	c = newsapi.NewClient(os.Args[2], newsapi.WithHTTPClient(http.DefaultClient))
}

func main() {
	commands := map[string]func(cid string, m string, args []string){
		"taco": Reply,
		"xkcd": Xkcd,
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
		input := strings.Split(m.Content, " ")
		if r, ok := commands[input[0]]; ok {
			r(m.ChannelID, m.Content, input[1:len(input)])
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
