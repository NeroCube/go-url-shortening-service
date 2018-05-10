package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
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
