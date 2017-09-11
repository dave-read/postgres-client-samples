package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/lib/pq"
)

var HOST, DATABASE, USER, PASSWORD, URI string

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	if cfenv.IsRunningOnCF() {
		appEnv, appErr := cfenv.Current()
		checkError(appErr)
		serviceName := getEnvVarOrExit("db-service")
		service, serviceError := appEnv.Services.WithName(serviceName)
		checkError(serviceError)
		credentials := service.Credentials
		HOST = credentials["hostname"].(string)
		DATABASE = credentials["name"].(string)
		USER = credentials["username"].(string)
		PASSWORD = credentials["password"].(string)
		URI = credentials["uri"].(string)

	} else {
		HOST = getEnvVarOrExit("hostname")
		DATABASE = getEnvVarOrExit("database")
		USER = getEnvVarOrExit("username")
		PASSWORD = getEnvVarOrExit("password")
		URI = getEnvVarOrExit("uri")
	}

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
	var catalog string
	var schema string
	var name string

	sqlStatement := "SELECT table_catalog, table_schema, table_name FROM information_schema.tables where table_type = 'BASE TABLE' and table_schema = 'information_schema';"
	rows, err := db.Query(sqlStatement)
	checkError(err)

	for rows.Next() {
		switch err := rows.Scan(&catalog, &schema, &name); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%s, %s, %s)\n", catalog, schema, name)
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
