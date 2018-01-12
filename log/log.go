package log

import (
	"log"
	"os"
)

var (
	Error = log.New(os.Stderr, "myboard ERR: ", log.LstdFlags)
	Info  = log.New(os.Stdout, "myboard INFO: ", log.LstdFlags)
)
