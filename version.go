package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	// get environment variables for MySql user and password
	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")

	// no user
	if user == "" {
		fmt.Println("Set SQL_USER environment variable")
	}

	// no password
	if password == "" {
		fmt.Println("Set SQL_PASSWORD environment variable")
	}

	// exit if no user and password
	if (user == "") || (password == "") {
		return
	}

	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password

	dsn := config.FormatDSN()
	fmt.Println(dsn)

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping to confirm connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connect and check the server version
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to:", version)
}
