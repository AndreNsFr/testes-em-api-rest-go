package main

import (
	"fmt"
	"net/http"
	"crud-produtos/controller"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao analisar formulário", http.StatusInternalServerError)
		return
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Dados recebidos:\n")
	fmt.Fprintf(w, "Nome: %s\n", nome)
	fmt.Fprintf(w, "Email: %s\n", email)
}

func main() {
	http.HandleFunc("/socorro", controllers.CriarProduto)
	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
