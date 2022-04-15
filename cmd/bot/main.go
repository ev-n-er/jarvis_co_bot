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

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8888"
	}

	if botServer, err := server.CreateNew(port, apiKey, serverUrl); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Starting listner at port %s", port)
		err = botServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}

}
