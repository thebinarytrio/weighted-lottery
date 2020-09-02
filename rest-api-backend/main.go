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

	// connecting to the database
	db, err = sql.Open("mysql", "root:@/lottery")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// creating new gorilla/mux router
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", getUsers).Methods("GET")

	// creating connection to local browser
	http.ListenAndServe(":30000", router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

// retrieves all users from the users table and response with JSON formatted variables
func getUsers(w http.ResponseWriter, r *http.Request) {

	// sets content-type to applications/json
	w.Header().Set("Content-Type", "application/json")

	var users []Users

	// query returns the rows of the database
	result, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// loop over result and for every iteration, creates new User instance
	// then scans the result into user object
	for result.Next() {
		var user Users
		err := result.Scan(&user.Userid, &user.Username)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	// encodes users array into JSON string
	json.NewEncoder(w).Encode(users)
	fmt.Println("Endpoint Hit: getUsers")
}
