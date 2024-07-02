package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathToURL struct {
	path string
	url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := pathsToUrls[r.URL.Path]
		if url != "" {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (yamlHandler http.HandlerFunc, err error) {
	parsedYaml, err := parseYAML(yml)

	if err != nil {
		return
	}

	builtMap := make(map[string]string)

	for _, singePathToUrl := range parsedYaml {
		builtMap[singePathToUrl.path] += singePathToUrl.url
	}

	yamlHandler = MapHandler(builtMap, fallback)
	return
}

func parseYAML(yamlData []byte) (pathsToURLs []pathToURL, err error) {
	err = yaml.Unmarshal(yamlData, &pathsToURLs)
	return
}
