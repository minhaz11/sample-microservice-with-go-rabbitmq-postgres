package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/minhaz11/authentication-service/data"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	//TODO connect to database DB

	//setup the config
	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
	  log.Panic(err)
	}
}