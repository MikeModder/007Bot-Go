package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func registerCommands() {
	handler.AddCommand("about", "Show information about the bot", false, false, 0, aboutCmd)
	handler.AddCommand("ping", "Check the bot's ping", false, false, 0, pingCmd)
	handler.AddCommand("version", "Show the versions of things", false, false, 0, versionCmd)

	// Add the default help command
	handler.AddDefaultHelpCommand()
}

func aboutCmd(ctx anpan.Context, _ []string) {
	embed := &discordgo.MessageEmbed{
		Title:       "About 007Bot (Golang edition):",
		Description: "Hello, I am [007Bot Golang Edition](https://github.com/MikeModder/007Bot-Go)! I am a rewrite of the original 007Bot in [Golang](https://golang.org) using the [anpan](https://github.com/MikeModder/anpan) command handler!",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Wow, look at me!",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}

func pingCmd(ctx anpan.Context, _ []string) {
	ping, _ := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Let me check that for you!")

	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne)

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf("Ping took `%s`!", took.String()),
	}

	ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, ping.ID, embed)
}

func versionCmd(ctx anpan.Context, _ []string) {
	embed := &discordgo.MessageEmbed{
		Title:       "007Bot Version(s)",
		Description: fmt.Sprintf("Golang: %s", runtime.Version()),
	}

	ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
