package main

import (
	"fmt"
  "net/http"
  "eventos"
)

func handler(w http.ResponseWriter, r *http.Request) {
    resultado, erros := eventos.ObtemEventos()
    fmt.Fprintf(w, "%v \nErros: %v", resultado, erros)
}

func main() {
    http.HandleFunc("/novembro-negro", handler)
    http.ListenAndServe(":8080", nil)
}
