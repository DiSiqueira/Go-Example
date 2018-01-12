package board

import (
	"fmt"

	"github.com/go-gorp/gorp"
)

type (
	// Board describes a board db entry
	Board struct {
		ID   uint64 `db:"id"`
		Text string `db:"text"`
	}

	BoardFinder interface {
		// Find returns a Board by ID
		Find(ID uint64) (*Board, error)
	}

	BoardUpdater interface {
		// Update updates a Board entry
		Update(b *Board) error
	}

	BoardInserter interface {
		// Insert inserts a Board into db
		Insert(b *Board) error
	}

	boardManager struct {
		dbMap gorp.SqlExecutor
	}
)

// NewBoardFinder inits and returns an instance of BoardFinder
func NewBoardFinder(dbMap gorp.SqlExecutor) BoardFinder {
	return newBoardManager(dbMap)
}

// NewBoardUpdater inits and returns an instance of BoardUpdater
func NewBoardUpdater(dbMap gorp.SqlExecutor) BoardUpdater {
	return newBoardManager(dbMap)
}

// NewBoardInserter inits and returns an instance of BoardInserter
func NewBoardInserter(dbMap gorp.SqlExecutor) BoardInserter {
	return newBoardManager(dbMap)
}

func newBoardManager(dbMap gorp.SqlExecutor) *boardManager {
	return &boardManager{dbMap}
}

func (m *boardManager) Find(ID uint64) (*Board, error) {
	var b Board

	if err := m.dbMap.SelectOne(&b, "SELECT id, name FROM board WHERE id = ?", ID); err != nil {
		return nil, fmt.Errorf("boardManager.Find: %s", err)
	}
	return &b, nil
}

func (m *boardManager) Update(b *Board) error {
	_, err := m.dbMap.Update(b)
	if err != nil {
		return fmt.Errorf("boardManager.Update: %s", err)
	}
	return nil
}

func (m *boardManager) Insert(b *Board) error {
	if err := m.dbMap.Insert(b); err != nil {
		return fmt.Errorf("boardManager.Insert: %s", err)
	}
	return nil
}
