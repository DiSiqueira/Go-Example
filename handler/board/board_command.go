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
	// PostBoardHandler implements
	// http.Handler interface
	// and serves GET board requests
	PostBoardHandler struct {
		finder      db.BoardFinder
		paramReader handler.URLParamReader
	}
)

// NewBoardQuery inits and returns an instance
// of GetBoardHandler
func NewBoardCommand(
	finder db.BoardFinder,
	paramReader handler.URLParamReader,
) http.Handler {
	return &PostBoardHandler{finder, paramReader}
}

// ServeHTTP implements http.Handler interface
func (h *PostBoardHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}