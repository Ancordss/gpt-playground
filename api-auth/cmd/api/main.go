package main

import (
	"api-auth/internal/auth"
	"api-auth/internal/server"
	"fmt"
)

func main() {
	auth.NewAuth()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
