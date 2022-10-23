package main

import (
	"calcium/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {

	args := os.Args[1:]

	port, err := utils.GetArgValue(args, "-port")

	if err != nil {
		fmt.Printf("error: Couldn't get value for parameter '-port' at args %v\n", args)
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	router.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("you are in home"))
	})

	log.Fatal(http.ListenAndServe(":"+port, router))

}
