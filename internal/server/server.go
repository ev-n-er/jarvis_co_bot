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
	"github.com/ev-n-er/jarvis_co_bot/internal/commands/tomorrow"
	"github.com/ev-n-er/jarvis_co_bot/internal/pkg/message"
)

type CommandHandler func(*message.Update) (*message.ResponseMessage, error)
type CallbackHandler func(*message.Update, *url.Values) (*message.EditMessage, error)

type BotServer struct {
	botUrl          string
	port            string
	responderUrl    string
	commandHandlers map[string]CommandHandler
	callbacks       map[string]CallbackHandler
}

func CreateNew(port string, apiKey string, responderUrl string) (*BotServer, error) {
	return &BotServer{
		botUrl:       fmt.Sprintf("https://api.telegram.org/bot%s/", apiKey),
		port:         port,
		responderUrl: responderUrl,
		commandHandlers: map[string]CommandHandler{
			"/start":    start.Handler,
			"/tomorrow": tomorrow.Handler,
			"/help":     help.Handler,
		},
		callbacks: map[string]CallbackHandler{
			"tomorrow": tomorrow.Callback,
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

	http.HandleFunc("/", server.handleRequest)
	http.ListenAndServe(server.port, nil)

	return nil
}

func (server *BotServer) handleRequest(w http.ResponseWriter, req *http.Request) {
	log.Printf("Got request %s", req.RequestURI)

	var update message.Update

	rb, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Req body: %v", string(rb))

	if err := json.Unmarshal(rb, &update); err == nil {
		log.Printf("JSON parsed: %#v", update)

		if update.CallbackQuery != nil {
			server.handleCallback(&update)
		} else {
			server.handleCommand(&update)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		log.Printf("Failed to parse JSON: %s", err)
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("I'm a teapot"))
	}
}

func (server *BotServer) handleCommand(update *message.Update) {
	if handler, ok := server.commandHandlers[update.Message.Text]; ok {
		if msg, err := handler(update); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Using %s responder", update.Message.Text)
			server.respond(msg)
		}
	} else {
		log.Print("No suitable responder")
	}
}

func (server *BotServer) handleCallback(update *message.Update) {

	if args, err := url.ParseQuery(update.CallbackQuery.Data); err == nil {

		command := args.Get("cmd")

		if handler, ok := server.callbacks[command]; ok {
			if msg, err := handler(update, &args); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Using %s responder for callback", command)
				server.edit(msg)
			}
		} else {
			log.Print("No suitable responder for callback")
		}
	} else {
		log.Fatalf("Could not parse command args: %v", err)
	}

}

func (server *BotServer) respond(m *message.ResponseMessage) error {

	url := fmt.Sprintf("%ssendMessage", server.botUrl)

	log.Printf("Will send: %#v to %s", *m, url)

	body, err := json.Marshal(*m)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(b))

	return nil
}

func (server *BotServer) edit(m *message.EditMessage) error {

	url := fmt.Sprintf("%seditMessageText", server.botUrl)

	log.Printf("Will send: %#v to %s", *m, url)

	body, err := json.Marshal(*m)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(b))

	return nil
}
