package help

import (
	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.ResponseMessage, error) {
	inMessage := (*update).Message
	return &message.ResponseMessage{
		Text:   "I can assist with organizing coworking visits.",
		ChatId: inMessage.Chat.Id,
	}, nil
}
