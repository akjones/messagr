package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func helloHandler(message string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("REQUEST")
		fmt.Println(r)

		fmt.Println("RESPONSE")
		fmt.Println(w)

		io.WriteString(w, message)
	}
}

func config_file_path() string {
	path := ""
	env := os.Environ()
	for _, val := range env {
		if strings.Contains(val, "MESSAGR_APP_CONFIG") {
			path = strings.Split(val, "=")[1]
			break
		}
	}

	if path == "" {
		path = "../config/production.yaml"
	}

	return path
}

func main() {
	the_message := message(config_file_path())
	fmt.Println("Messagr app starting...")
	fmt.Println("Config file: ", config_file_path())
	fmt.Println("Message: ", the_message)

	handler := helloHandler(the_message)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
