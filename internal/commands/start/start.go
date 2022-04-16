package start

import (
	"fmt"

	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.Message, error) {
	inMessage := (*update).Message
	var text string
	if inMessage.From.FirstName != "" {
		text = inMessage.From.FirstName
	} else {
		text = inMessage.From.Username
	}

	return &message.Message{
		Text: fmt.Sprintf("Hey %s", text),
		Chat: message.Chat{
			Id: inMessage.Chat.Id,
		},
	}, nil
}
