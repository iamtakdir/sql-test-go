package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
