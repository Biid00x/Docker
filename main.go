package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	msg := "Hello, World!"
	if envMsg := os.Getenv("MESSAGE"); envMsg != "" {
		msg = envMsg
	}
	addr := ":8080"
	if envPortStr := os.Getenv("PORT"); envPortStr != "" {
		port, err := strconv.Atoi(envPortStr)
		if err != nil {
			log.Fatalf("invalid port from env var: %v", err)
		}
		addr = fmt.Sprintf(":%d", port)
	}

	responseBody := fmt.Sprintf(`{"message": "%s"}`, msg)
	log.Printf("MOTD API listening on [%s]", addr)
	log.Printf("Message is [%s]", msg)
	http.ListenAndServe(addr,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(responseBody))
				log.Printf("%v %v", r.Method, r.URL.Path)
			}))
}