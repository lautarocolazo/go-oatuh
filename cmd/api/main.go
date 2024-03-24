package main

import (
	"fmt"

	"go-oauth/internal/auth"
	"go-oauth/internal/server"
)

func main() {
	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
