package main

import (
	db "mydatabase"
)

func main() {
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8000")
}
