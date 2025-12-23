package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token          string // Discord bot token
	Log_Channel_ID string
	Server_ID      string
	cmdPrefix      string = "rp."
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	Log_Channel_ID = os.Getenv("LOG_CHANNEL_ID")
	Server_ID = os.Getenv("SERVER_ID")
}

func main() {

	//make new session
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	r, w := io.Pipe()

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	/*
	   if Server_ID != dg.Guild.ID {
	       dg.Close()
	       return
	   }
	*/
	dg.ChannelMessageSend(Log_Channel_ID, "Bot Started!")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.ChannelMessageSend(Log_Channel_ID, "Bot Bot Stopping!")
	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	fmt.Println("test 1")
	if m.Author.ID == s.State.User.ID || m.GuildID != Server_ID {
		return
	}

	fmt.Println(m.Content)
	fmt.Println(cmdPrefix)

	cmd, hasPref := strings.CutPrefix(m.Content, cmdPrefix)

	if !hasPref {
		return
	}

	fmt.Println("wow you did it")
	return

	args := []string{m.Author.ID}

	exit := false
	for !exit {
		var (
			arg   string
			found bool
		)

		arg, cmd, found = strings.Cut(cmd, " ")
		if !found {
			exit = true
			break
		}
		args = append(args, arg)

	}

}
