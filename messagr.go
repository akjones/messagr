package main

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
)

func message(path string) string {
	configFileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return "whoops"
	}

	config := make(map[string]string)
	yamlErr := yaml.Unmarshal([]byte(configFileContent), &config)

	if yamlErr != nil {
		return "double whoops"
	}

	return config["the_message"]
}

func hello(w http.ResponseWriter, r *http.Request) {
	the_message := message("../config/production.yaml")

	io.WriteString(w, the_message)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8000", nil)
}
