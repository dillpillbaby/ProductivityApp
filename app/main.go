package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234ASDFjkl;"
	dbname   = "ProductivityApp"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	test_db, err := sql.Open("postgres", psqlInfo)
	check_db_error(err)

	
	err = test_db.Ping()
	check_db_error(err)

	create_test_entry("Dylan", test_db)
	create_test_entry("Eli", test_db)
	create_test_entry("Alex", test_db)

	getAll := `SELECT name from test`
	items, e := test_db.Query(getAll)
	check_db_error(e)
	defer items.Close()
	for items.Next() {
		var name string
		if err := items.Scan(&name); err != nil {
				panic(err)
		}
		fmt.Println(name)
	}
	if err := items.Err(); err != nil {
		panic(err)
	}
	check_db_error(e)
	defer test_db.Close()
}

func check_db_error(err error) {
    if err != nil {
        panic(err)
    }
}
func create_test_entry(name string, db *sql.DB){
	insrtTest := `INSERT INTO test("name") values('dylan')`
	_, e := db.Exec(insrtTest)
	check_db_error(e)
}