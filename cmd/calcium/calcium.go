package main

import (
	"calcium/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, `web\public`))

	FileServer(router, "/files", filesDir)
	router.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("you are in home"))
	})

	log.Fatal(http.ListenAndServe(":"+port, router))

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
