package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/root-of-a-tree/remote_exec_bot/pkg/config"
)

const guildID = "810610920313061428"

func main() {
	config, err := config.InitConfig()
	if err != nil {
		fmt.Println("error initializing bot configuration,", err)
		return
	}
	fmt.Println("Initialized config.")
	fmt.Printf("Servers: %+v\n", config.Servers)
	fmt.Printf("Scripts: %+v\n", config.Scripts)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// TODO: create application command
	// commands := []*discordgo.ApplicationCommand{
	// 	{
	// 		Name:        "example",
	// 		Description: "description",
	// 	},
	// 	{
	// 		Name:        "options",
	// 		Description: "example of options",
	// 		Options: []*discordgo.ApplicationCommandOption{
	// 			{
	// 				Type:        discordgo.ApplicationCommandOptionString,
	// 				Name:        "string-option",
	// 				Description: "string option",
	// 				Required:    true,
	// 			},
	// 		},
	// 	},
	// }

	// TODO: create a handler for each application command
	// commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// 	"example": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 			Data: &discordgo.InteractionApplicationCommandResponseData{
	// 				Content: "This is an example slash command",
	// 			},
	// 		})
	// 	},
	// 	"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 		stringOpt := i.Data.Options[0].IntValue()
	// 		msgformat :=
	// 			` Option entered:
	// 		> string-option: %s
	// `
	// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 			Data: &discordgo.InteractionApplicationCommandResponseData{
	// 				Content: fmt.Sprintf(msgformat, stringOpt),
	// 			},
	// 		})
	// 	},
	// }

	// TODO: add all the handlers using discordgo.InteractionCreate as the event handler type
	// dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 	if handler, ok := commandHandlers[i.Data.Name]; ok {
	// 		handler(s, i)
	// 	}
	// })

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Rex is ready!")
	})

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// TODO: create all our application commands
	// for _, v := range commands {
	// 	_, err := dg.ApplicationCommandCreate(appID, guildID, v)
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	// 	}
	// }

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Gracefully shutting down")
	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
