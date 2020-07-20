package main

import (
	"go/src/go-api/db"
	"go/src/go-api/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
