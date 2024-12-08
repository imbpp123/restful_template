package main

import (
	"log"
	"net/http"

	"app/internal"
)

func main() {
	if err := http.ListenAndServe(":3333", internal.RouterInitializer()); err != nil {
		log.Fatal(err)
	}
}
