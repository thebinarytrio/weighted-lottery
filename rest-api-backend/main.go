package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}

	defer db.Close()

	// creating new gorilla/mux router
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{userid}", getUser).Methods("GET")
	router.HandleFunc("/users/{userid}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{userid}", deleteUser).Methods("DELETE")

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
		log.Fatal(err)
	}

	defer result.Close()

	// loop over result and for every iteration, creates new User instance
	// then scans the result into user object
	for result.Next() {
		var user Users
		err := result.Scan(&user.Userid, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// encodes users array into JSON string
	json.NewEncoder(w).Encode(users)
	fmt.Println("Endpoint Hit: getUsers")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := db.Prepare("INSERT INTO `users`(username) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	username := keyVal["username"]

	_, err = stmt.Exec(username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "New post was created")
	fmt.Println("Endpoint Hit: createUser")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT userid, username FROM `users` WHERE userid = ?", params["userid"])
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	var user Users

	for result.Next() {
		err := result.Scan(&user.Userid, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
	}

	json.NewEncoder(w).Encode(user)
	fmt.Println("Endpoint Hit: getUser")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE `users` SET username = ? WHERE userid = ?")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newUser := keyVal["username"]

	_, err = stmt.Exec(newUser, params["userid"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "User with ID = %s was updated", params["userid"])
	fmt.Println("Endpoint Hit: updateUser")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM `users` WHERE userid = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(params["userid"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "User with ID =%s was deleted", params["userid"])
	fmt.Println("Endpoint Hit: deleteUser")
}
