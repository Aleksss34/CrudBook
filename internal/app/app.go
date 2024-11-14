package app

import (
	"CrudBook/internal/config"
	handlers "CrudBook/internal/transport/http"
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	rout := handlers.Handlers()

	if err := http.ListenAndServe(config.PORT, rout); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server open")

}
