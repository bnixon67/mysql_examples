package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// dsn is provided as first command line argument
	if len(os.Args) != 2 {
		log.Fatal("Must provide DSN as first command line argument")
	}
	dsn := os.Args[1]

	// create and populate config structure for DSN to login
	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// update dsn from config structure
	dsn = config.FormatDSN()
	fmt.Println("DSN:", dsn)

	// Create the database handle
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

	fmt.Println("Version:", version)
}
