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
}

func main() {
	fmt.Println("> Starting up bot...")

	dg, err := discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		fmt.Printf("[error] failed to create a session: %v\n", err)
		os.Exit(1)
	}

	handler = anpan.NewCommandHandler(botConfig.Prefix, botConfig.Owners, true, true)
	handler.StatusHandler.SetSwitchInterval(botConfig.StatusInterval)
	handler.StatusHandler.SetEntries(botConfig.Statuses)
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
