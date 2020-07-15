package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

func main() {
	connectDB()
}

func connectDB() {

	DBconnStr := "host=aio1 port=5432 user=gpadmin dbname=gpadmin password=abc123 sslmode=disable"
	db, err := sql.Open("postgres", DBconnStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(db))
}
