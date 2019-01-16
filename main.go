package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

var (
	botConfig Config
	handler   anpan.CommandHandler

	GitCommit, GitBranch, Version string
)

func init() {
	fmt.Println("> Loading and parsing config...")

	cfg, err := os.Open("config.json")
	if err != nil {
		fmt.Sprintf("[error] failed to open config.json: %v\n", err)
		os.Exit(1)
	}

	err = json.NewDecoder(cfg).Decode(&botConfig)
	if err != nil {
		fmt.Sprintf("[error] failed to decode config.json: %v\n", err)
		os.Exit(1)
	}

	if GitBranch == "" {
		GitBranch = "unknown"
		GitCommit = "abcd123"
		Version = "unknown"
	}
}

func main() {
	fmt.Printf("[007Bot-Go Version %s (%s-%s)]\n", Version, GitCommit, GitBranch)
	fmt.Println("> Starting up bot...")

	dg, err := discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		fmt.Printf("[error] failed to create a session: %v\n", err)
		os.Exit(1)
	}

	handler = anpan.NewCommandHandler(botConfig.Prefix, botConfig.Owners, true, true)
	handler.StatusHandler.SetSwitchInterval(botConfig.StatusInterval)
	handler.StatusHandler.SetEntries(botConfig.Statuses)
	handler.SetPrerunFunc(beforeOnMessage)
	//handler.SetDebug(true)

	registerCommands()

	dg.AddHandler(handler.OnMessage)
	dg.AddHandler(handler.StatusHandler.OnReady)

	err = dg.Open()
	if err != nil {
		fmt.Printf("[error] failed to open a session: %v\n", err)
		os.Exit(1)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("> Shutting down...")
	dg.Close()
}

func beforeOnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, has := AFKUsers[m.Author.ID]
	if has {
		delete(AFKUsers, m.Author.ID)

		embed := &discordgo.MessageEmbed{
			Title:       "Welcome back!",
			Description: fmt.Sprintf("Welcome back, *%s*! Your AFK status has been removed.", m.Author.String()),
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	for i := 0; i < len(m.Mentions); i++ {
		mention := m.Mentions[i]
		afkEntry, has := AFKUsers[mention.ID]
		if has {
			embed := &discordgo.MessageEmbed{
				Title:       fmt.Sprintf("*%s* is AFK!", mention.String()),
				Description: fmt.Sprintf("*%s* is currently AFK!\n\n```%s```", mention.String(), afkEntry.Message),
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}
	}
}
