package main

import (
	"fmt"
	"github.com/okikechinonso/cmd/server"
)

func main() {
	fmt.Println("hello")
	s := server.NewServer()
	s.Start()
}
