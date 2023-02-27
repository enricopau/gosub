package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/enricopau/gosub/database"
	"github.com/enricopau/gosub/server"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	client, err, disconnect := database.Connect()
	defer disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mail, err := database.Mail(ctx, "67267@test.test", database.MailListCollection(client))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mail)

	glue, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Handler:      glue.Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
