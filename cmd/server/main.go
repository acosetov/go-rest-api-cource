package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/acosetov/go-rest-api-cource/internal/transport/http"
)

// App - the struct witch contains thinks like pointers
// to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed a set up server")
		return err
	}

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
