package main

import (
	"github.com/ikbarfp/bumder/internal/app"
	"log"
)

func main() {
	srv := app.New()
	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}
}
