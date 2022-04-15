package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/ev-n-er/jarvis_co_bot/internal/commands/help"
	"github.com/ev-n-er/jarvis_co_bot/internal/commands/start"
	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

type CommandHandler func(*message.Update) (*message.Message, error)

type BotServer struct {
	botUrl          string
	port            string
	responderUrl    string
	commandHandlers map[string]CommandHandler
}

func CreateNew(port string, apiKey string, responderUrl string) (*BotServer, error) {
	return &BotServer{
		botUrl:       fmt.Sprintf("https://api.telegram.org/bot%s/", apiKey),
		port:         port,
		responderUrl: responderUrl,
		commandHandlers: map[string]CommandHandler{
			"/start": start.Handler,
			"/help":  help.Handler,
		},
	}, nil
}

func (server *BotServer) ListenAndServe() error {

	params := url.Values{}
	params.Add("url", server.responderUrl)

	resp, err := http.Get(fmt.Sprintf("%ssetWebhook?%s", server.botUrl, params.Encode()))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		responseText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to set webhook: %s", err)
		}

		return fmt.Errorf("Failed to set webhook: %s", responseText)
	} else {
		log.Print("Webhook set, listening...")
	}

	http.HandleFunc("*", server.handleRequest)
	http.ListenAndServe(server.port, nil)

	return nil
}

func (server *BotServer) handleRequest(w http.ResponseWriter, req *http.Request) {

	var update message.Update

	if err := json.NewDecoder(req.Body).Decode(&update); err != nil {
		log.Fatal(err)
		return
	}

	if handler, ok := server.commandHandlers[update.Message.Text]; ok {
		if msg, err := handler(&update); err != nil {
			log.Fatal(err)
		} else {
			server.respond(msg)
		}
	}

}

func (server *BotServer) respond(m *message.Message) error {

	url := fmt.Sprintf("%ssendMessage", server.botUrl)
	body, err := json.Marshal(*m)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
