package main

import "log"

const webPort = "80"

type Application struct {
}

func main() {
	app := &Application{}

	if err := app.routes().Run(":" + webPort); err != nil {
		log.Fatalf("Failed to start broker service: %s", err)
	}
}
