package main

import (
	"github.com/antonioofdz/personalprojectdra/pkg/handlers"
)

func main() {
	/*if err := database.InitDB(); err != nil {
		panic(err)
	}*/

	handlers.LoadRoutes()
}
