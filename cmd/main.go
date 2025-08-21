package main

import (
	"log"
	"whatsapp-api/internal/app"
	"whatsapp-api/util"
)

func main() {
	cfg, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
