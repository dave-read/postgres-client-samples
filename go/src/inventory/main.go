package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var (
	// Initialize connection constants.
	HOST     = getEnvVarOrExit("hostname")
	DATABASE = getEnvVarOrExit("database")
	USER     = getEnvVarOrExit("username")
	PASSWORD = getEnvVarOrExit("password")
	URI      = getEnvVarOrExit("uri")
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Initialize connection string from properties
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
	// Initialize connection object.
	fmt.Printf("Connecting using connection string: %s\n", connectionString)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)
	queryDatabase(db)

	// use URI directly
	fmt.Printf("Connecting using uri: %s\n", URI)
	db, err = sql.Open("postgres", URI)
	checkError(err)
	queryDatabase(db)

	// parse properties out of the URI
	parsedUrl, err := url.Parse(URI)
	checkError(err)
	hostName := parsedUrl.Hostname()
	userName := parsedUrl.User.Username()
	password, _ := parsedUrl.User.Password()
	database := strings.TrimPrefix(parsedUrl.Path, "/")
	parsedConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", hostName, userName, password, database)
	fmt.Printf("Connecting using connection string: %s\n", parsedConnectionString)
	db, err = sql.Open("postgres", parsedConnectionString)
	checkError(err)
	queryDatabase(db)

}

func queryDatabase(db *sql.DB) {

	err := db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	// Read rows from table.
	var id int
	var name string
	var quantity int

	sql_statement := "SELECT * from inventory;"
	rows, err := db.Query(sql_statement)
	checkError(err)

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}
	fmt.Println("")
}

// getEnvVarOrExit returns the value of specified environment variable or terminates if it's not defined.
func getEnvVarOrExit(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		fmt.Printf("Missing environment variable %s\n", varName)
		os.Exit(1)
	}

	return value
}
