package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const HTTP_PORT = "3333"

func getRoot(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "api root response")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	fmt.Printf("starting api server on port %s", HTTP_PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%s", HTTP_PORT), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
