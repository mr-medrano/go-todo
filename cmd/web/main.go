package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"examples.mrmedano.todo/internal/database"
)

type config struct {
	addr string
	dsn  string
}

type Application struct{}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "postgres://admin:password@localhost:5432/postgres", "PSQL data source name")

	flag.Parse()

	app := &Application{}

	pgPool, err := database.NewPGXPool(context.Background(), cfg.dsn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer pgPool.Close()

	srv := &http.Server{
		Addr:         cfg.addr,
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	fmt.Println(err.Error())
	os.Exit(1)

}
