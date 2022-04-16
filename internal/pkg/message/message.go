package message

type Update struct {
	Id      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
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
