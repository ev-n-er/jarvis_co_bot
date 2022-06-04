package message

type Update struct {
	Id            int            `json:"update_id"`
	Message       Message        `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type Message struct {
	MessageId      int          `json:"message_id"`
	Text           string       `json:"text"`
	From           User         `json:"from"`
	Chat           Chat         `json:"chat"`
	ReplyToMessage *MessageBase `json:"reply_to_message"`
}

type MessageBase struct {
	Text string `json:"text"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
}

type ResponseMessage struct {
	Text             string                `json:"text"`
	ChatId           int                   `json:"chat_id"`
	ReplyMarkup      *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ReplyToMessageId int                   `json:"reply_to_message_id"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	Id int `json:"id"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageId string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
}

type CallbackAnswer struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
}

type EditMessage struct {
	ChatId          int                   `json:"chat_id"`
	MessageId       int                   `json:"message_id"`
	InlineMessageId string                `json:"inline_message_id"`
	Text            string                `json:"text"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`
}
