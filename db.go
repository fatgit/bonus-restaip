package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func dbConnect() *sql.DB {

	if 	dbPort := os.Getenv("DB_PORT"); len(dbPort) == 0 {
		dbPort = "3306"
	}
	dataSource := os.Getenv("DB_ADMIN") + ":" + os.Getenv("DB_ADMIN_PASS") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (c Client) createDatabase() {
	db := dbConnect()
	defer db.Close()

	_,err := db.Exec("CREATE DATABASE "+ c.Database + ";")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CREATE DATABASE ", c.Database)
	_,err = db.Exec("GRANT USAGE ON *.* to " + "'" + c.Name + "'@'%' identified by '" + c.clientMysqlPassword() + "';")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CREATE USER ", c.Name)

	_,err = db.Exec("GRANT ALL PRIVILEGES ON " + c.Database + ".* to " + "'" + c.Name + "'@'%';")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GRANT PERMISSON ", c.Name, " on ", c.Database)
}

func (c Client) dropDatabase()  {
	db := dbConnect()
	defer db.Close()

	_,err := db.Exec("DROP DATABASE "+ c.Database + ";")
	if err != nil {
		log.Fatal(err)
	}
}
