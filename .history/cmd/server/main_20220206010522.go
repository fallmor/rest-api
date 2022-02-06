package main

import (
	"fmt"
)

type App struct {
}

func (ap *App) Run() error {
	fmt.Println("Starting the server")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	return nil
}
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Can't start the server")
	}

}
