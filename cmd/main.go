package main

import (
	"log"

	goo "github.com/SHCDevelops/go_first"
	"github.com/SHCDevelops/go_first/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(goo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
