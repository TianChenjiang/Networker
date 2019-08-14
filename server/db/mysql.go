package db

import "database/sql"
import "log"

var My *sql.DB // MySQL

func init() {
	var err error
	username := "root"
	password := "mysql"
	host := "localhost"
	port := "33006"
	dataname := "citicup"
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	//Set up sql connection
	My, err = sql.Open("mysql", dns)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = My.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
