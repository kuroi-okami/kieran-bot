package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string

func main() {

	if token == "" {
		fmt.Println("No token provided.")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Running is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateStatus(0, "Doing Robo-God's work")

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	rand.Seed(time.Now().UnixNano())

	if m.Content == "!story" {
		storyId := rand.Intn(10)
		switch storyId {
		case 1:{
			s.ChannelMessageSend(m.ChannelID, "I once fed a homeless man crumbs")
		}
		case 2:{
			s.ChannelMessageSend(m.ChannelID, "Once James and Will froze my shoes")
		}
		case 3:{
			s.ChannelMessageSend(m.ChannelID, "I love jesus")
		}
		case 4:{
			s.ChannelMessageSend(m.ChannelID, "Have you meet jesus?")
		}
		case 5:{
			s.ChannelMessageSend(m.ChannelID, "I have cute feet")
		}
		case 6:{
			s.ChannelMessageSend(m.ChannelID, "What is love?")
		}
		case 7:{
			s.ChannelMessageSend(m.ChannelID, "You guys are fudged in the head")
		}
		case 8:{
			s.ChannelMessageSend(m.ChannelID, "I ate a cake that was too terribly good to eat badly")
		}
		case 9:{
			s.ChannelMessageSend(m.ChannelID, "I miss limelight girl")
		}
		case 10:{
			s.ChannelMessageSend(m.ChannelID, "You guys are lame")
		}
		default:{
			s.ChannelMessageSend(m.ChannelID, "I am all out of stories")
		}
		}
	}
