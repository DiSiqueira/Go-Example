package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/disiqueira/Go-Example/log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/disiqueira/Go-Example/db/board"
	"github.com/disiqueira/Go-Example/db/eventsource"
)

const maxConnectAttempts = 5

// NewDB returns
func NewDB(dsn string, maxOpen, maxIdle, attempt int) (*gorp.DbMap, error) {
	errfmt := "NewDB: %s"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf(errfmt, err)
	}

	if err := conn.Ping(); err != nil {
		if attempt < maxConnectAttempts {
			log.Error.Printf(errfmt, err)
			time.Sleep(time.Second)
			return NewDB(dsn, maxOpen, maxIdle, attempt+1)
		}

		return nil, fmt.Errorf(errfmt, err)
	}

	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	dbMap := &gorp.DbMap{
		Db: conn,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	// set mappings
	dbMap.AddTableWithName(board.Board{}, "board").SetKeys(true, "ID")
	dbMap.AddTableWithName(eventsource.EventSource{}, "eventsource").SetKeys(true, "ID")

	return dbMap, nil
}
