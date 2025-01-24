package main

import (
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

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
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
