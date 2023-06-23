package main

import "database/sql"
import "authentication/data"

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	print("Starting authentication service...")
}