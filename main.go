package main

import (
	"fmt"
	"log"

	"github.com/jestape/hackovid-dyb-api/src/app"
	"github.com/jestape/hackovid-dyb-api/src/config"
)

func main() {
	config, err := config.GetConfig()

	if err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	app := &app.App{}
	app.Initialize(config)

	log.Printf("Serving application at PORT: %d", config.Port)
	app.Run(fmt.Sprintf(":%d", config.Port))
}
