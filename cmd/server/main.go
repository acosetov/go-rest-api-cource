package main

import "fmt"

// App - the struct witch contains thinks like pointers
// to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
