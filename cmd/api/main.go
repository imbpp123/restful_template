package main

import (
	"log"
	"net/http"

	"app/internal"
)

func main() {
	articleHandler := internal.InitializeArticleHandler()
	if err := http.ListenAndServe(":3333", internal.RouterInitializer(articleHandler)); err != nil {
		log.Fatal(err)
	}
}
