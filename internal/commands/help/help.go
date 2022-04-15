package help

import (
	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.Message, error) {
	inMessage := (*update).Message
	return &message.Message{
		Text: "I can assist with organizing coworking visits.",
		Chat: message.Chat{
			Id: inMessage.Chat.Id,
		},
	}, nil
}
