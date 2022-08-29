package main

import (
	"log"
	"net/http"
	controller "servers_distribute/Controller"
)

func main() {
	log.Println("Starting Servers_distribute...")
	controller.InitRoutes()
	http.ListenAndServeTLS(":8080", "./server.crt", "./server.key", controller.GetHandler())
}
