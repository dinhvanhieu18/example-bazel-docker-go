package main

import (
	"log"
	"net/http"
	"time"

	"demo/app/internal"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hi", internal.HiHandler).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Println("Start serving ...")
	log.Fatal(srv.ListenAndServe())
}
