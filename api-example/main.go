package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	ApiVersion    string    `json:"api_version"`
}

func main() {
	port, envVarIsSet := os.LookupEnv("PORT")
	if !envVarIsSet {
		slog.Info("PORT env var was not set. Using 8080 as default.")
		port = "8080"
	}

	http.HandleFunc("/api/ping", func(w http.ResponseWriter, req *http.Request) {
    slog.Info("ping received")
		json_response, _ := json.Marshal(response{
			Message: "pong",
			Code:    http.StatusOK,
      ApiVersion: "v3",
		})

		w.Write(json_response)
	})
	
  http.HandleFunc("/api/pong", func(w http.ResponseWriter, req *http.Request) {
    slog.Info("pong received")
    json_response, _ := json.Marshal(response{
			Message: "ping",
			Code:    http.StatusOK,
      ApiVersion: "v3",
		})

		w.Write(json_response)
	})
	
	slog.Info(fmt.Sprintf("Listening on: %s \n", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
