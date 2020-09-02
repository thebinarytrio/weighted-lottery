package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Users struct contains 2 properies that represents all the users on the site
type Users struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@/lottery")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":30000", router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []Users

	result, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var user Users
		err := result.Scan(&user.Userid, &user.Username)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
	fmt.Println("Endpoint Hit: getUsers")
}
