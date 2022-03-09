package main

import (
	"fmt"
	"net/http"

	"github.com/fallmor/rest-api/internal/comment"
	"github.com/fallmor/rest-api/internal/database"
	myTransport "github.com/fallmor/rest-api/internal/transport/http"
)

type App struct {
}

func (ap *App) Run() error {
	fmt.Println("Starting the server")

	var err error
	db, err := database.DatabaseSetup()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.Newcomment(db)
	handler := myTransport.NewRouter(commentService)
	handler.SetupRoutes()
	fmt.Println("connected to the database")
	if err := http.ListenAndServe(":8080", handler.Route); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Can't start the server")
	}

}
