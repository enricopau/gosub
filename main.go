package main

import (
	"context"
	"time"

	"github.com/enricopau/gosub/database"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	go checkRemoveableMails(ctx)
	<-ctx.Done()
	// g := server.NewGlue(ctx, "127.0.0.1:8000")
	// g.StartServer()
}

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
	err = database.DeleteUnconfirmedEntriesAfter(ctx, time.Now().Add(-12*time.Hour), database.MailListCollection(client))
	return err
}
