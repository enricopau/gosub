package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enricopau/gosub/database"
)

func main() {
	client, err, disconnect := database.Connect()
	defer disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	entries, err := database.Confirmed(database.MailListCollection(client))
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Printf("Mail: %v", entry.Mail)
	}

	// http.HandleFunc("/", Server)
	// http.ListenAndServe(":8080", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
