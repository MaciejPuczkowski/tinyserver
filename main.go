package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)


var root = flag.String("r", "/www", "root of the files")
var addr = flag.String("a", ":80", "address to listen")

func handleError(err error, writer http.ResponseWriter) {
	if err == os.ErrNotExist {
		writer.WriteHeader(404)
		return
	}
	if err != nil {
		writer.WriteHeader(500)
		return
	}
}

func main() {
	flag.Parse()
	log.Printf("starting server, listening on %s with root: %s", *addr, *root)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			var path string
			if request.URL.Path == "/" {
				path = fmt.Sprintf("%s/%s", *root, "index.html")
			} else {
				path = fmt.Sprintf("%s/%s", *root, strings.TrimLeft(request.URL.Path, "/"))
			}
			content, err := os.ReadFile(path)
			handleError(err, writer)
			if err != nil {
				log.Printf("ERROR: %v", err)
				return
			}
			writer.WriteHeader(200)
			writer.Write(content)
		}
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
