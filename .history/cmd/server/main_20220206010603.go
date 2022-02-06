package main

import (
	"fmt"
	"net/http"

	myTransport "github.com/fallmor/rest-api/internal/transport/http"
)

type App struct {
}

func (ap *App) Run() error {
	fmt.Println("Starting the server")

	handler := myTransport.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
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
