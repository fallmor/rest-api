package main

import (
	"fmt"

	myTransport "github.com/fallmor/rest-api/internal/transport/http"
)

type App struct {
}

func (ap *App) Run() error {
	fmt.Println("Starting the server")

	router := myTransport.NewRouter()
	return nil
}
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Can't start the server")
	}

}
