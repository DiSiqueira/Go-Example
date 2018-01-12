package board

import (
	"encoding/json"
	"fmt"

	"github.com/disiqueira/Go-Example/db/board"
)

type (
	// Board describes a board API model
	Board struct {
		ID      uint64 `json:"id"`
		Text    string `json:"text"`
		SelfURL string `json:"self_url"`
	}
)

func (b *Board) MarshalJSON() ([]byte, error) {
	b.SelfURL = fmt.Sprintf("https://myboard.io/boards/%d", b.ID)
	type alias Board

	return json.Marshal(alias(*b))
}

func (b *Board) fromDB(dbB *board.Board) {
	b.ID = dbB.ID
	b.Text = dbB.Text
}
