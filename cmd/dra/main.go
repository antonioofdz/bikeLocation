package main

import (
	"github.com/antonioofdz/personalprojectdra/pkg/database"
	"github.com/antonioofdz/personalprojectdra/pkg/handlers"
)

func init() {
	if err := database.InitDB(); err != nil {
		panic(err)
	}
}

func main() {
	handlers.LoadRoutes()
}
