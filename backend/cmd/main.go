package main

import (
	"database/sql"
	"fmt"

	// mySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Open up the database connection
	db, err := sql.Open("mysql", "root:@/lottery")

	// if there was an error opening the connection, print connection failed
	if err != nil {
		fmt.Println("Connection Failed")
	}

	defer db.Close()

	var (
		userid   int
		username string
	)

	row := db.QueryRow("SELECT userid, username FROM users WHERE userid = ?", 1)
	err = row.Scan(&userid, &username)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userid)
	fmt.Println(username)
}
