package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Produto struct {
    Nome  string  `json:"nome"`
    Email string `json:"email"`
	Telefone string `json:"telefone"`
}

func CriarProduto(w http.ResponseWriter, r *http.Request) {
    var struct_produto Produto

    err := json.NewDecoder(r.Body).Decode(&struct_produto)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }


	fmt.Fprint(w,struct_produto)
	
    
}