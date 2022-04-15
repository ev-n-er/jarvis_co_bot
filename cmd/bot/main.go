package main

import (
	"log"
	"os"

	"github.com/ev-n-er/jarvis_co_bot/internal/server"
)

func main() {

	log.Print("Starting jarvis")

	apiKey := os.Getenv("JARVIS_BOT_API_KEY")
	serverUrl := os.Getenv("JARVIS_BOT_SERVER_URL")

	if botServer, err := server.CreateNew(":8888", apiKey, serverUrl); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Starting listner")
		err = botServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}

}
