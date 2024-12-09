package main

import (
	"fmt"
	"log"
	"net/http"

	"app/internal"
)

func main() {
	fmt.Print("Lets do the job!")
	articleHandler := internal.InitializeArticleHandler()
	fmt.Print("Handler was created")
	if err := http.ListenAndServe(":3333", internal.RouterInitializer(articleHandler)); err != nil {
		log.Fatal(err)
	}
}
