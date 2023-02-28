package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/enricopau/gosub/database"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	client, dc, err := database.Connect()
	defer dc(ctx)
	if err != nil {
		log.Fatal(err)
	}
	twelve := time.Now().Add(-12 * time.Hour)
	fmt.Printf("Time: %s\n", twelve)
	err = database.DeleteEntriesAfter(ctx, twelve, database.MailListCollection(client))
	if err != nil {
		log.Fatal(err)
	}

	// g := server.NewGlue(ctx, "127.0.0.1:8000")
	// g.StartServer()
}

func RemoveCheck() {
	// make this like 10 hours in PROD
	t := time.NewTicker(12 * time.Second)

	for {
		select {
		case <-t.C:
			removeAfter(ctx)
		}
	}
}

func removeAfter(ctx context.Context) error {
	client, dc, err := database.Connect()
	defer dc(ctx)
	if err != nil {
		return err
	}
	err = database.DeleteEntriesAfter(ctx, time.Now().Add(-12*time.Hour), database.MailListCollection(client))
	return err
}
