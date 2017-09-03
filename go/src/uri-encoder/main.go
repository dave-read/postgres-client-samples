package main

import (
	"fmt"
	"net/url"
	"os"
)

const (
	service = ".postgres.database.azure.com"
	scheme  = "postgres"
	port    = "5432"
)

var (
	server   = getEnvVarOrExit("server")
	database = getEnvVarOrExit("database")
	username = getEnvVarOrExit("user")
	password = getEnvVarOrExit("password")
)

// postgres://test$1@dr-test:{your_password}@dr-test.postgres.database.azure.com:5432/test$1
func main() {

	uri := url.URL{}
	uri.Scheme = scheme
	uri.Host = server + service + ":" + port
	uri.User = url.UserPassword(username+"@"+server, password)
	uri.Path = database
	// removed query string as there's no single portable representation and its not in the bind sample
	//q := uri.Query()
	//q.Set("ssl", "true")
	//uri.RawQuery = q.Encode()
	fmt.Println(uri.String())
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
