package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	h := func(w http.ResponseWriter, _ *http.Request) {
		user := os.Getenv("DB_USERNAME")
		pass := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/appname", user, pass, host))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var name string
		err = db.QueryRow("SELECT name FROM users LIMIT 1").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "Hello, %s!", name)
	}

	http.HandleFunc("/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
