package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/horlabyc/gophercises/URL-SHORTNER/urlshort"
)

const PORT = "8080"

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received; %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := defaultMux()

	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "http://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	fileName := flag.String("file-name", "redirect.yaml", "Yaml file containing the redirect url mappings")
	flag.Parse()

	fileExtension := filepath.Ext(*fileName)

	dataBytes, err := urlshort.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	var handler http.Handler
	var error error
	if fileExtension == ".yaml" {
		fmt.Printf("Processing %s file extension\n", fileExtension)
		handler, error = urlshort.YAMLHandler(dataBytes, mapHandler)
		if error != nil {
			panic(error)
		}
	} else if fileExtension == ".json" {
		fmt.Printf("Processing %s file extension\n", fileExtension)
		handler, error = urlshort.JSONHandler(dataBytes, mapHandler)
		if error != nil {
			panic(error)
		}
	} else {
		log.Fatal("Allowed file extension are: .json, .yaml")
	}

	fmt.Println("Starting server on port" + PORT)
	http.ListenAndServe("localhost:8080", LoggerMiddleware(handler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
