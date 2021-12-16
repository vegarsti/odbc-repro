package main

import (
	"database/sql"
	"fmt"

	_ "github.com/vegarsti/odbc" // odbc database driver
)

func main() {
	dsc := `...`
	db, err := sql.Open("odbc", dsc)
	if err != nil {
		fmt.Printf("sql open: %v\n", err)
		return
	}
	defer db.Close()
	var s string
	if err := db.QueryRow("select 'hello world'").Scan(&s); err != nil {
		fmt.Printf("query: %v\n", err)
		return
	}
	fmt.Printf("got '%s'\n", s)
}
