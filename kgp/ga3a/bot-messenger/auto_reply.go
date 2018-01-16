package botMessenger

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	."./models"
)

func ProcessMessage(event Messaging) {
	log.Println("Recive message!")
	log.Println(event.Sender.ID)
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: Message{
			Text: "Hello!",
		},
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	url := fmt.Sprintf(FACEBOOK_API, os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	log.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}