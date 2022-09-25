package api

import (
	TransactionService "Challenge/internal/repositories/transactions"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {

	transactions, err := TransactionService.TransactionRepo.GetAllTransactions()
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

func Create(w http.ResponseWriter, r *http.Request) {

	trans := TransactionService.Transaction{}
	json.NewDecoder(r.Body).Decode(&trans)
	_, err := TransactionService.TransactionRepo.CreateTransaction(&trans)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't Add Transaction %s", err), 200)
		return
	}
	w.Write([]byte("The transaction has been added \n"))
	json.NewEncoder(w).Encode(trans)
	w.WriteHeader(200)
}

func HandleRequest() {
	Router := chi.NewRouter()
	Router.Get("/transactions", GetAllTransactions)
	Router.Post("/transactions/create", Create)
	log.Fatal(http.ListenAndServe(":8080", Router))
}
