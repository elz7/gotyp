package main

import (
	"io"
	"log"
)

func initDebugConsole(w io.Writer) {
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.SetOutput(w)
}
