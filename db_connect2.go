package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sf "github.com/snowflakedb/gosnowflake"
)

var db *sql.DB

func main() {
	conn, err := GetConnection()
	if err != nil {
		log.Fatal("Error while connection ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Successfully get the connection...")
}

func GetConnection() (conn *sql.Conn, err error) {

	//	https://qg83307.ap-southeast-1.snowflakecomputing.com // linkedin //

	dns, err := sf.DSN(&sf.Config{
		Account:  "qg83307",
		User:     "rajuugupta",
		Password: "Raju@0987",
		Database: "demo",
		Schema:   "demo_schema",
		Region:   "ap-southeast-1",
	})

	if err != nil {
		log.Fatal("Error while DNS string: ", err)
		return conn, err
	}

	fmt.Println("raju dsn: ", dns)
	db, err := sql.Open("snowflake", dns)
	if err != nil {
		log.Fatal("Error while open DB: ", err)
		return conn, err
	}
	defer db.Close()

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(3 * time.Second)
	log.Println("Database Connection Successful..!")

	conn, err = db.Conn(context.Background())
	query_string := "CREATE ROLE db_hr_ab"
	// res, err := conn.QueryContext(context.Background(), query_string)

	res, err := conn.ExecContext(context.Background(), query_string)

	if err != nil {
		log.Fatal("Error while open connection: ", err)
		return conn, err
	}

	log.Println(fmt.Sprintf("result: %s", res))
	log.Println("Connection Successful..!")
	return conn, err
}
