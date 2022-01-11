package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.1.109"
	port     = 5432
	user     = "postgres"
	password = "Krebsisgarb"
	dbname   = "recipe"
)

func Connect() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func Create_User(email string, password string){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}


	sqlStatement := `
	insert into users (email, password)
	values ($1, $2)`
	_, err = db.Exec(sqlStatement, email, password)
	if err != nil {
		panic(err)
	}
	
}