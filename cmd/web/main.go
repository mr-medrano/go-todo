package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Application struct{}

func main() {
	app := &Application{}

	srv := &http.Server{
		Addr:         ":9000",
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	fmt.Println(err.Error())
	os.Exit(1)

}
