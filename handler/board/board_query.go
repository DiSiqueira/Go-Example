package board

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/disiqueira/Go-Example/db/board"
	"github.com/disiqueira/Go-Example/handler"
	"github.com/disiqueira/Go-Example/log"
)

type (
	// GetBoardHandler implements
	// http.Handler interface
	// and serves GET board requests
	GetBoardHandler struct {
		finder      db.BoardFinder
		paramReader handler.URLParamReader
	}
)

// NewBoardQuery inits and returns an instance
// of GetBoardHandler
func NewBoardQuery(
	finder db.BoardFinder,
	paramReader handler.URLParamReader,
) http.Handler {
	return &GetBoardHandler{finder, paramReader}
}

// ServeHTTP implements http.Handler interface
func (h *GetBoardHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	errfmt := "GetBoardHandler: %s"

	boardIDStr := h.paramReader.Read(req, "boardID")
	boardID, err := strconv.ParseUint(boardIDStr, 10, 64)
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "invalid board ID provided", http.StatusBadRequest)
		return
	}

	dbBoard, err := h.finder.Find(boardID)
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "board not found", http.StatusNotFound)
		return
	}

	board := &Board{}
	board.fromDB(dbBoard)

	b, err := json.Marshal([]*Board{board})
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
