package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	DB_USER     := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME     := os.Getenv("POSTGRES_DB")
)

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println(dbinfo)
	checkErr(err)
	defer db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
