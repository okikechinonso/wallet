package main

import (
	"github.com/okikechinonso/cmd/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
