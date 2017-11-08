package main

import (
	"log"

	"github.com/ECAllen/continuous_improvement/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
