package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/enricopau/gosub/database"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go shutdown(cancel)

	go checkRemoveableMails(ctx)
	<-ctx.Done()

	// g := server.NewGlue(ctx, "127.0.0.1:8000")
	// g.StartServer()
}

// checkRemoveableMails checks for mails that are "unconfirmed" for 24 hours and deletes them automatically.
func checkRemoveableMails(ctx context.Context) {
	// make this like 10 hours in PROD
	t := time.NewTicker(12 * time.Second)

outerloop:
	for {
		select {
		case <-t.C:
			removeAfter(ctx)
		case <-ctx.Done():
			break outerloop
		}
	}
}

func removeAfter(ctx context.Context) error {
	client, dc, err := database.Connect()
	defer dc(ctx)
	if err != nil {
		return err
	}
	err = database.DeleteUnconfirmedEntriesAfter(ctx, time.Now().Add(-24*time.Hour), database.MailListCollection(client))
	return err
}

func shutdown(cancel context.CancelFunc) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
	fmt.Println("shutting down")
	cancel()
}
