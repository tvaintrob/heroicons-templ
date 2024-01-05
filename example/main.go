package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = ExamplePage().Render(r.Context(), w)
	})

  log.Println("Server running at http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
