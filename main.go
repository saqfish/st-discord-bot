package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

func init() {
	var err error

	session, err = discordgo.New("Bot " + os.Args[1])

	if len(os.Args) < 3 {
		log.Fatal("go run main.go [DISCORD KEY] [NEWS-API KEY]")
		os.Exit(1)
	}

	if err != nil || len(os.Args) < 3 {
		log.Fatal("New bot error")
		os.Exit(1)
	}
}

func main() {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		input := strings.Split(m.Content, " ")
		if r, ok := Commands[input[0]]; ok {
			r(m.ChannelID, input[1:]...)
		}
	})

	err := session.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}
