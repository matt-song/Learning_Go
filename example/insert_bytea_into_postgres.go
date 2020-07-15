package main

import (
	"database/sql"
	"io/ioutil"
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

	/* Create the table */
	println("Creating the table...\n")
	db.Exec("drop table if exists test_binary_table;")
	db.Exec("create table test_binary_table (id int, file bytea);")

	/* Read the content of video(binary) file */
	println("Reading the content of binary file...\n")
	content, err := ioutil.ReadFile("/Users/xsong/Downloads/sample-mp4-file.mp4")
	if err != nil {
		log.Fatal(err)
	}
	bytes := []byte(content)

	/* Insert into the table */
	println("Inserting the content into table...\n")
	db.Exec("insert into test_binary_table values(1, $1::bytea)", bytes)
	rows, err := db.Query("select file from test_binary_table where id=1")

	/* Write the file back to disk */
	for rows.Next() {
		var output []byte
		rows.Scan(&output)
		println("Wrting the data to disk...\n")
		err := ioutil.WriteFile("/tmp/output.mp4", output, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}
