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
	initiallyCreate()
	create_star_systems()
	create_planets()
	create_belts()
	populate()
	getAllBelts()
	dropTable("belts", 5)
	getAllBelts()
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

func populate() {
	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close();
	_, err = db.Exec("USE " + db_name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("INSERT INTO belts VALUES (5, 3)")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added Data to Table Belts")
	}
}

	func getAllBelts() {
		fmt.Println("Printing belts")
		var (
			id int
			ressources string
		)
		db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close();
		_,err = db.Exec("USE "+db_name)
		if err != nil {
			fmt.Println(err)
		}
		rows, err := db.Query("select * from belts")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		fmt.Println("Closed rows")
		for rows.Next() {
			err := rows.Scan(&id, &ressources)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id, ressources)
		}
		err = rows.Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	/**
	Should only be used for debugging, handle with care
	 */
func dropTable(tableName string, id int) () {
	fmt.Printf("Dropping table %s\n", tableName)
	db, err := sql.Open("mysql", "moritz:moritz@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close();
	_, err = db.Exec("USE " + db_name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("DElETE FROM " + tableName + " WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Dropped Table %s\n", tableName)
	}
}
