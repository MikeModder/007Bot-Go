package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func registerCommands() {
	handler.AddCommand("about", "Show information about the bot", false, 0, aboutCmd)
	handler.AddCommand("ping", "Check the bot's ping", false, 0, pingCmd)
}

func aboutCmd(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	embed := &discordgo.MessageEmbed{
		Title:       "About 007Bot (Golang edition):",
		Description: "Hello, I am [007Bot Golang Edition](https://github.com/MikeModder/007Bot-Go)! I am a rewrite of the original 007Bot in [Golang](https://golang.org) using the [anpan](https://github.com/MikeModder/anpan) command handler!",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Wow, look at me!",
			IconURL: s.State.User.AvatarURL("512"),
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func pingCmd(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	ping, _ := s.ChannelMessageSend(m.ChannelID, "Pong...")

	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne)

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf("Ping took `%s`!", took.String()),
	}

	s.ChannelMessageEditEmbed(m.ChannelID, ping.ID, embed)
}
