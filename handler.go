package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	GetResponse(msg string) string
}

func (babur *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Log messages
	if len(m.Message.Content) > 120 {
		fmt.Printf("MSG: %s ...\n", m.Message.Content[:120])
	} else {
		fmt.Printf("MSG: %s\n", m.Message.Content)
	}

	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, handler := range babur.handlers {
		if content := handler.GetResponse(m.Message.Content); content != "" {
			sendMessage(s, m, content)
		}
	}
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	// Mention the user in the first line
	msg := fmt.Sprintf("<@%s> \n%s", m.Message.Author.ID, content)

	// Log reponses
	fmt.Printf("RESPONSE: %s\n", msg)

	if _, err := s.ChannelMessageSend(m.ChannelID, msg); err != nil {
		fmt.Println("ERROR: Send message: ", err)
		return
	}
}
