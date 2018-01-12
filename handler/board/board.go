package board

import (
	"encoding/json"
	"github.com/disiqueira/Go-Example/db/board"
	"net/http"
)

type (
	// Board describes a board API model
	Board struct {
		ID   uint64 `json:"id"`
		Text string `json:"text"`
	}
)

func (b *Board) MarshalJSON() ([]byte, error) {
	type alias Board

	return json.Marshal(alias(*b))
}

func (b *Board) fromDB(dbB *board.Board) {
	b.ID = dbB.ID
	b.Text = dbB.Text
}

func (b *Board) Bind(r *http.Request) error {
	return nil
}
