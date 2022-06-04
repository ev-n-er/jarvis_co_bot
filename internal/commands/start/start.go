package start

import (
	"fmt"

	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.ResponseMessage, error) {
	inMessage := (*update).Message
	var text string
	if inMessage.From.FirstName != "" {
		text = inMessage.From.FirstName
	} else {
		text = inMessage.From.Username
	}

	return &message.ResponseMessage{
		Text:        fmt.Sprintf("Hey %s", text),
		ChatId:      inMessage.Chat.Id,
		ReplyMarkup: nil,
	}, nil
}
