package board

import (
	"encoding/json"
	"fmt"
	"github.com/disiqueira/Go-Example/db/events"
	"github.com/disiqueira/Go-Example/db/eventsource"
	"github.com/disiqueira/Go-Example/handler"
	"github.com/go-chi/render"
	"net/http"
)

type (
	// PostBoardHandler implements
	// http.Handler interface
	// and serves GET board requests
	PostBoardHandler struct {
		inserter    eventsource.EventSourceInserter
		paramReader handler.URLParamReader
	}
)

// NewBoardQuery inits and returns an instance
// of GetBoardHandler
func NewBoardCommand(
	inserter eventsource.EventSourceInserter,
	paramReader handler.URLParamReader,
) http.Handler {
	return &PostBoardHandler{inserter, paramReader}
}

// ServeHTTP implements http.Handler interface
func (h *PostBoardHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	data := &Board{}
	if err := render.Bind(req, data); err != nil {
		http.Error(w, "invalid board Object provided", http.StatusBadRequest)
		return
	}

	bwc := events.BoardWasCreated{Board: fmt.Sprintf("%v", data)}
	bwcSTR, err := json.Marshal(bwc)
	if err != nil {
		http.Error(w, "invalid board Object provided when generating JSON", http.StatusBadRequest)
		return
	}

	event := eventsource.EventSource{}
	event.Event = string(bwcSTR)

	err = h.inserter.Insert(&event)
	if err != nil {
		http.Error(w, "error when inserting into the event database", http.StatusInternalServerError)
		return
	}

	fmt.Println(data)
	fmt.Println(event)
	fmt.Println(bwcSTR)

	w.WriteHeader(http.StatusCreated)
}
