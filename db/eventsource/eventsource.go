package eventsource

import (
	"fmt"
	"github.com/go-gorp/gorp"
)

type (
	EventSource struct {
		ID    uint64 `db:"id"`
		Event string `db:"event"`
	}

	EventSourceInserter interface {
		// Insert inserts a Board into db
		Insert(es *EventSource) error
	}

	EventSourceManager struct {
		dbMap gorp.SqlExecutor
	}
)

func newEventSourceManager(dbMap gorp.SqlExecutor) *EventSourceManager {
	return &EventSourceManager{dbMap}
}

func NewEventSourceInserter(dbMap gorp.SqlExecutor) EventSourceInserter {
	return newEventSourceManager(dbMap)
}

func (esm *EventSourceManager) Insert(es *EventSource) error {
	if err := esm.dbMap.Insert(es); err != nil {
		return fmt.Errorf("EventSourceManager.Insert: %s", err)
	}
	return nil
}
