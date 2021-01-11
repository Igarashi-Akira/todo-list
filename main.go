package main

import (
	"github.com/Igarashi-Akira/todo-list/db"
	"github.com/Igarashi-Akira/todo-list/server"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}