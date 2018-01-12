package main

import (
	"github.com/go-chi/chi"
	"fmt"
	"github.com/disiqueira/Go-Example/config"
	"github.com/disiqueira/Go-Example/log"
	"github.com/disiqueira/Go-Example/db"
	"github.com/disiqueira/Go-Example/handler"
	"github.com/disiqueira/Go-Example/handler/board"
	dbb "github.com/disiqueira/Go-Example/db/board"
	"net/http"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error.Fatal(err)
	}

	dbMap, err := db.NewDB(cfg.MySQLDSN, 5, 5, 0)
	if err != nil {
		log.Error.Fatal(err)
	}

	boardFinder := dbb.NewBoardFinder(dbMap)
	urlReader := handler.NewURLParamReader()

	router := chi.NewRouter()
	router.Route("/boards", func(r chi.Router) {
		r.Get("/{boardID:[0-9]+}", board.NewBoardQuery(boardFinder, urlReader).ServeHTTP)
		r.Post("/", board.NewBoardCommand(boardFinder, urlReader).ServeHTTP)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Error.Fatal(err)
	}
}
