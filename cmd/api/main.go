package main

import (
	"fmt"
	"github.com/goexpert/desafio-tecnico-stress-test/internal/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
