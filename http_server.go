/*
Copyright 2020 Jonathan Gonzalez V.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ""

	if os.Getenv("HTTP_PORT") == "" {
		port = ":8080"
	} else {
		port = os.Getenv("HTTP_PORT")
	}

	http.HandleFunc("/", handler)
	log.Printf("Listening on %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Not supported\n")
		return
	}

	path := ""
	if os.Getenv("HTTP_PATH") == "" {
		path, _ = os.Getwd()
	} else {
		path = os.Getenv("HTTP_PATH")
	}
	http.ServeFile(w, r, path+r.URL.Path)
}
