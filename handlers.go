package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	store *Store
}

func NewAPI(store *Store) *API {
	return &API{store: store}
}

// POST /accounts
func (a *API) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.DocumentNumber == "" {
		http.Error(w, "document number is required", http.StatusBadRequest)
		return
	}
	acc := a.store.CreateAccount(req.DocumentNumber)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(acc)

}

// GET /accounts/{id}
func (a *API) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["accountId"]
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}
	acc, err := a.store.GetAcount(id)
	if err != nil {
		http.Error(w, "account not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(acc)
}

// POST /transactions
func (a *API) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//basic validation
	if req.AccountID == 0 || req.OperationTypeID == 0 {
		http.Error(w, "account_id and operation_type_id are required", http.StatusBadRequest)
		return
	}

	t, err := a.store.CreateTransaction(req)
	if err != nil {
		switch err {
		case ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case ErrInvalidOperation:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case ErrInvalidAmountSign:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(t)
}
