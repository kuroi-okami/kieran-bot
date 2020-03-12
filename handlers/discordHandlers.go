package handlers

import (
	"github.com/bwmarrin/discordgo"
	"kieran-bot/handlers/delegates"
)

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateStatus(0, "Doing Robo-God's work")

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Delegate handling
	delegates.KieranSpeak(s, m)
	delegates.RemindClark(s, m)
	delegates.LoopSteve(s, m)
	delegates.DdosRob(s, m)
}
