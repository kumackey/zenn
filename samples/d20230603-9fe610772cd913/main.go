package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		println(err)
	}
	defer sqldb.Close()

	// 疎通確認
	if err := sqldb.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping error: %v\n", err)
		os.Exit(1)
	}

	println("OK")
}
