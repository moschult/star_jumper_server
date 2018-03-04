package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
var (
	db_name string = "star_jump"
)

func main() {
	initiallyCreate();
	create_star_systems();
	create_planets();
	create_belts();
}

func initiallyCreate() (*sql.DB) {

	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("CREATE DATABASE "+db_name)
	if err != nil {
		fmt.Println(err)
	}

	_,err = db.Exec("USE "+db_name)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func create_star_systems() (){
	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close();

	_,err = db.Exec("USE "+db_name)
	if err != nil {
		fmt.Println(err)
	}

	_,err = db.Exec("CREATE TABLE star_systems ( id INTEGER, planets INTEGER, belts INTEGER )")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully created table star_systems")
	}
}

func create_planets() (){
	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close();

	_,err = db.Exec("USE "+db_name)
	if err != nil {
		fmt.Println(err)
	}

	_,err = db.Exec("CREATE TABLE planets ( id INTEGER, population INTEGER, farmland INTEGER )")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully created table planets")
	}
}

func create_belts() (){
	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close();

	_,err = db.Exec("USE "+db_name)
	if err != nil {
		fmt.Println(err)
	}

	_,err = db.Exec("CREATE TABLE belts ( id INTEGER, resources INTEGER )")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully created table belts")
	}

}

func populate()  {

}