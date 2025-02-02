package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados")
		panic(err)
	}

	db.AutoMigrate(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	r.ParseForm()
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	db.Create(&user)

	fmt.Fprintf(w, "Usuário criado: %+v", user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {

	http.HandleFunc("/users/create", createUser)

	http.HandleFunc("/users", getUsers)

	fmt.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
