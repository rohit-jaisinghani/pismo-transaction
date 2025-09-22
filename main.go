package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	store := NewStore()
	api := NewAPI(store)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", api.CreateAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountId}", api.GetAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", api.CreateTransactionHandler).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal((http.ListenAndServe(addr, r)))
}
