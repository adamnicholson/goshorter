package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/adamnicholson/goshorter/shortner"
)

func main() {
	app := shortner.Container{}
	app.Boot()

	server := shortner.HttpServer{App: app}
	server.Listen("8080")
}
