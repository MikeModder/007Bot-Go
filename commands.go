package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

var (
	AFKUsers = make(map[string]AFKEntry)
)

func registerCommands() {
	handler.AddCommand("about", "Show information about the bot", false, false, 0, aboutCmd)
	handler.AddCommand("ping", "Check the bot's ping", false, false, 0, pingCmd)
	handler.AddCommand("version", "Show the versions of things", false, false, 0, versionCmd)
	handler.AddCommand("afk", "Go afk", false, false, 0, afkCmd)

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

	ctx.ReplyEmbed(embed)
}

func pingCmd(ctx anpan.Context, _ []string) {
	ping, _ := ctx.Reply("Let me check that for you!")

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

	ctx.ReplyEmbed(embed)
}

func afkCmd(ctx anpan.Context, args []string) {
	msg := strings.Join(args, " ")
	if msg == "" {
		msg = "No message set."
	} else if len(msg) > 500 {
		embed := &discordgo.MessageEmbed{
			Title:       "Error!",
			Description: fmt.Sprintf("Sorry *%s*, your afk message cannot be longer than 500 characters!", ctx.User.String()),
			Color:       0xff0000,
		}

		ctx.ReplyEmbed(embed)
		return
	}

	AFKUsers[ctx.User.ID] = AFKEntry{
		Message: msg,
		Set:     time.Now(),
	}

	embed := &discordgo.MessageEmbed{
		Title:       "See you later!",
		Description: fmt.Sprintf("*%s*, you have been marked as afk!\n\nWhen your friends mention you, they'll see the message you left (if any). Don't worry about removing your afk status, I'll remove it as soon as you send a message.", ctx.User.String()),
		//Timestamp:   string(time.Now().Unix()),
	}

	ctx.ReplyEmbed(embed)
	return
}
