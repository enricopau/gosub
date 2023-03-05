package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/enricopau/gosub/database"
	"github.com/enricopau/gosub/token"
	"github.com/gorilla/mux"
)

type Glue struct {
	srv http.Server
}

func NewGlue(ctx context.Context, addr string) *Glue {
	r := mux.NewRouter()
	r.HandleFunc("/requesttoken/{mail}", RequestToken(ctx))
	http.Handle("/", r)

	return &Glue{
		srv: http.Server{
			Handler:      r,
			Addr:         addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (g *Glue) StartServer() {
	g.srv.ListenAndServe()
}

// RequestToken generates a new unconfirmed MailEntry and with a new token.
func RequestToken(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		client, dc, err := database.Connect()
		defer dc(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("error connecting to the database - %s", err), http.StatusInternalServerError)
			return
		}

		mail := mux.Vars(r)["mail"]
		newToken, err := token.GenerateToken()
		if err != nil {
			http.Error(w, fmt.Sprintf("error generating new token - %s", err), http.StatusInternalServerError)
			return
		}
		entry := database.NewMailEntry(mail, newToken, false)
		_, err = database.AddEntry(ctx, entry, database.MailListCollection(client))
		if err != nil {
			http.Error(w, fmt.Sprintf("error adding new mail to database - %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
