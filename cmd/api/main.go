package main

import (
	"fmt"
	"log"
	"net/http"

	"app/internal"
)

func main() {
	articleHandler := internal.InitializeArticleHandler()
	fmt.Print("Lets do the job!")
	if err := http.ListenAndServe(":3333", internal.RouterInitializer(articleHandler)); err != nil {
		log.Fatal(err)
	}
}
