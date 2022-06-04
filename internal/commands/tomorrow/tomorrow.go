package tomorrow

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ev-n-er/jarvis_co_bot/internal/db"
	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

func Handler(update *message.Update) (*message.ResponseMessage, error) {
	inMessage := (*update).Message
	return &message.ResponseMessage{
		Text:   "Will you be in the office tomorrow?",
		ChatId: inMessage.Chat.Id,
		ReplyMarkup: &message.InlineKeyboardMarkup{
			InlineKeyboard: [][]message.InlineKeyboardButton{
				{
					message.InlineKeyboardButton{Text: "Yes", CallbackData: "cmd=tomorrow&choice=1"},
				},
				{
					message.InlineKeyboardButton{Text: "Maybe", CallbackData: "cmd=tomorrow&choice=0"},
				},
				{
					message.InlineKeyboardButton{Text: "No", CallbackData: "cmd=tomorrow&choice=-1"},
				},
			},
		},
	}, nil
}

func Callback(update *message.Update, args *url.Values) (*message.EditMessage, error) {

	dbClient := db.Create()

	if dbClient == nil {
		return nil, fmt.Errorf("Could not create DB client")
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	choice, err := strconv.ParseInt(args.Get("choice"), 10, 8)
	if err != nil {
		return nil, fmt.Errorf("Could not identify choice")
	}

	created := dbClient.CreateVisit(uint((*update).CallbackQuery.From.Id), today, int8(choice))

	if !created {
		return nil, fmt.Errorf("Could not save answer")
	}

	inMessage := (*update).CallbackQuery.Message
	return &message.EditMessage{
		Text:      "Answer accepted",
		ChatId:    inMessage.Chat.Id,
		MessageId: inMessage.MessageId,
		ReplyMarkup: &message.InlineKeyboardMarkup{
			InlineKeyboard: [][]message.InlineKeyboardButton{},
		},
	}, nil
}
