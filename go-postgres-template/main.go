package main

import (
	"go-postgres/conf"
	"go-postgres/db"
	"go-postgres/logrus"
	"go-postgres/server"
)

func main() {

	logrus.InitLogrus()
	conf.LoadConfig()
	db.ConnectToDB()
	server.StartServer()

}
