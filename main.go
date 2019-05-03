package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"projetos/megarunning/prova"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/provas", prova.Handler)        //.Methods("GET")
	r.HandleFunc("/api/provas/{id}", prova.HandlerId) //.Methods("DELETE")
	r.HandleFunc("/api/provas/{id}/inscricao", prova.HandlerInscricao).Methods("POST")
	http.Handle("/api/", r)

	// http.HandleFunc("/api/provas")
	fs := http.Dir("public/")
	http.Handle("/", http.FileServer(fs))
	fmt.Println("Rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
