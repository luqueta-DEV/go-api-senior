package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello yatan")

}

func main() {

	http.HandleFunc("/", home)

	fmt.Println("server ta rodando lisin meu nobre na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error", err)
	}

}
