package start

import (
	"fmt"

	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.Message, error) {
	inMessage := (*update).Message
	var text string
	if inMessage.User.FirstName != "" {
		text = inMessage.User.FirstName
	} else {
		text = inMessage.User.Username
	}

	return &message.Message{
		Text: fmt.Sprintf("Hey %s", text),
		Chat: message.Chat{
			Id: inMessage.Chat.Id,
		},
	}, nil
}
