package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"examples.mrmedano.todo/config"
	"examples.mrmedano.todo/internal/database"
	"examples.mrmedano.todo/internal/models"
)

type Application struct {
	tasks *models.TaskModel
}

func main() {
	pgPool, err := database.NewDBPool(database.DatabaseConfig{
		Username: config.DB_USERNAME,
		Password: config.DB_PASSWORD,
		Hostname: config.DB_HOSTNAME,
		Port:     config.DB_PORT,
		DBName:   config.DB_NAME,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer pgPool.Close()

	app := &Application{
		tasks: &models.TaskModel{DB: pgPool},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.API_PORT),
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	fmt.Println(err.Error())
	os.Exit(1)

}
