package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// getEnvOrMessage retrieves the value of the environment variable,
// if present, named by the key, returning the value and true. If the
// key is not present, a message is printed and the returned value
// will be empty and the boolean will be false.
func getEnvOrMessage(key string) (result string, ok bool) {
	result, ok = os.LookupEnv(key)
	if !ok {
		fmt.Printf("Environment variable %s is not set.\n", key)
	}
	return
}

func main() {
	// get environment variables or print message if not set
	user, _ := getEnvOrMessage("SQL_USER")
	password, _ := getEnvOrMessage("SQL_PASSWORD")

	// exit if no user, or password
	if (user == "") || (password == "") {
		return
	}

	// create and populate config structure for DSN to login
	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password

	// create dsn from config structure
	dsn := config.FormatDSN()
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
