package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

const (
	service = ".postgres.database.azure.com"
	scheme  = "postgres"
	port    = 5432
)

var (
	username = `a"b`    //getEnvVarOrExit("user")
	password = `c"d`    //getEnvVarOrExit("password")
	database = "db"     //getEnvVarOrExit("database")
	server   = "server" //getEnvVarOrExit("server")
)

/*

 {
  "hostname": "db-hostname",
  "port": 5432,
  "database": "postgres",
  "username": "username",
  "password": "password",
  "uri": "postgres://username:password@db-hostname:5432/dbname",
  "uuid": "4ac6cb37-f486-4d71-a339-eb46bce4e399",
  "allocated_storage": 10,
  "maintenance_window": "fri:03:00-fri:03:30"
 }

*/

// ServiceDefinition is serialized to JSON
type ServiceDefinition struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	URI      string `json:"uri,omitempty"`
}

// postgres://test$1@dr-test:{your_password}@dr-test.postgres.database.azure.com:5432/test$1
func main() {

	uri := url.URL{}
	uri.Scheme = scheme
	uri.Host = server + service + ":" + string(port)
	uri.User = url.UserPassword(username+"@"+server, password)
	uri.Path = database
	// removed query string as there's no single portable representation and its not in the bind sample
	//q := uri.Query()
	//q.Set("ssl", "true")
	//uri.RawQuery = q.Encode()
	serviceDefintion := ServiceDefinition{
		Hostname: server,
		Port:     port,
		Database: database,
		Password: password,
		Username: uri.User.Username(),
		URI:      uri.String(),
	}

	fmt.Println(uri.String())

	fmt.Printf("struct:%+v", serviceDefintion)

	fmt.Println(marshalToJSON(serviceDefintion))

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

func encodeToJSON(v interface{}) string {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	err := enc.Encode(v)
	if err != nil {
		return "ERROR:" + err.Error()
	}
	return buf.String()
}

func marshalToJSON(v interface{}) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "ERROR:" + err.Error()
	}
	return string(j)
}
