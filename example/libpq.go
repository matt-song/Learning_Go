package main

// https://godoc.org/github.com/lib/pq
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	/* Connect to DB */
	connStr := "host=aio1 port=5432 user=gpadmin dbname=gpadmin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	/* query something */
	fmt.Printf("testing the query...\n")
	var queryResult []string
	rows, queryErr := db.Query("select distinct(host) from test_report;")
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	// Loop through the rows and store the table names.
	for rows.Next() {
		var host string
		err = rows.Scan(&host)
		if err != nil {
			fmt.Errorf("Error extracting the rows of the list of tables: %v", err)
		}
		// fmt.Printf(host + "\n")
		queryResult = append(queryResult, host)
	}
	for _, targetHost := range queryResult {
		fmt.Printf("find host: " + targetHost + "\n")
	}

}

/*
* dbname - The name of the database to connect to
* user - The user to sign in as
* password - The user's password
* host - The host to connect to. Values that start with / are for unix
  domain sockets. (default is localhost)
* port - The port to bind to. (default is 5432)
* sslmode - Whether or not to use SSL (default is require, this is not
  the default for libpq)
* fallback_application_name - An application_name to fall back to if one isn't provided.
* connect_timeout - Maximum wait for connection, in seconds. Zero or
  not specified means wait indefinitely.
* sslcert - Cert file location. The file must contain PEM encoded data.
* sslkey - Key file location. The file must contain PEM encoded data.
* sslrootcert - The location of the root certificate file. The file
  must contain PEM encoded data.

Valid values for sslmode are:

* disable - No SSL
* require - Always SSL (skip verification)
* verify-ca - Always SSL (verify that the certificate presented by the
  server was signed by a trusted CA)
* verify-full - Always SSL (verify that the certification presented by
  the server was signed by a trusted CA and the server host name
  matches the one in the certificate)

*/
