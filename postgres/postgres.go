package postgres

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Info() string {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	checkErr(err)

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_HOST"), port)

	return dbinfo
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
