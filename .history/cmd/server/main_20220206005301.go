package main

import "fmt"

type App struct {
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Can't start the server")
	}

}
