package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
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
	database, _ := getEnvOrMessage("SQL_DB")

	// exit if no user, password, or database
	/*
		if (user == "") || (password == "") || (database == "") {
			return
		}
	*/

	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password
	config.DBName = database

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

	var (
		host   string
		usr    string
		passwd string
	)

	rows, err := db.Query("select Host, User, Password from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cols)

	for rows.Next() {
		err := rows.Scan(&host, &usr, &passwd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(host, usr, passwd)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
