package main

import (
	"flag"
	"fmt"
	"net/http"

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

	yamlFileName := flag.String("yaml-file-name", "redirect.yaml", "Yaml file containing the redirect url mappings")
	flag.Parse()

	yamlDataBytes, err := urlshort.ReadFile(*yamlFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Yaml file name is " + *yamlFileName)

	yamlHandler, err := urlshort.YAMLHandler(yamlDataBytes, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on port" + PORT)
	http.ListenAndServe("localhost:8080", LoggerMiddleware(yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
