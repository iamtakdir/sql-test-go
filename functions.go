package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func get_all() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()

	results, err := db.Query("select * from customers ")

	if err != nil {
		log.Fatal("Error in fetching")
	}

	js_data := map[int]string{}
	for results.Next() {
		var (
			id   int
			name string
		)

		err = results.Scan(&id, &name)
		if err != nil {
			log.Fatal("Unable to fetch from row")
		}

		js_data[id] = name

	}

	json_data, _ := json.Marshal(js_data)

	fmt.Println(string(json_data))

}

func insert_into(key int, value string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Connection error")
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO `customers` (`id`, `name`) VALUES (?,?) ")
	if err != nil {
		log.Fatal("error in inserting")
	}
	insert.Exec(key, value)
	if err != nil {
		log.Fatal("Error in execution ")
	}

	fmt.Println("Inserted successfully")

	get_all()
}

func update_data(key int, value string) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Error in connection :Update")
	}
	defer db.Close()

	update, err := db.Prepare("UPDATE customers SET name = ? WHERE id = ?")
	fmt.Println(update)
	if err != nil {
		log.Fatal("error in updating ")
	}
	update.Exec(value, key)

	get_all()

}

func delete_data(id int) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Error in connection :Update")
	}
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM `customers` WHERE `customers`.`id` = ?")

	if err != nil {
		log.Fatal("error in updating ")
	}
	delete.Exec(id)

	get_all()

}

func marshalling(js map[int]string) {

	json.Marshal(js)
	fmt.Println(js)

	fmt.Println(js)
}
