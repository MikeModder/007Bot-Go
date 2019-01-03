package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func registerCommands() {
	handler.AddCommand("about", "Show information about the bot", false, 0, aboutCmd)
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
