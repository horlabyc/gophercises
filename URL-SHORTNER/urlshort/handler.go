package urlshort

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			log.Printf("Redirecting %s to %s\n", path, dest)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	// Convert YAML array to map
	pathsToUrls := buildPathMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildPathMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func parseYaml(yamlData []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlData, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
