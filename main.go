package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	connStr := "user=user password=password dbname=db host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")

	createTable(db)

	//insertUser(db, "Ivan", "ivan@gmail.com")

	getUserByID(db, 1)

	updateUserEmail(db, 1, "newivan@gmail.com")

	getUserByID(db, 1)

	deleteUser(db, 1)
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
}

func insertUser(db *sql.DB, name, email string) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	result, err := db.Exec(query, name, email)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User inserted. Rows affected: %d\n", rowsAffected)
}

func getUserByID(db *sql.DB, id int) {
	var name, email string
	query := "SELECT name, email FROM users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&name, &email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %s, Email: %s\n", name, email)
}

func updateUserEmail(db *sql.DB, id int, newEmail string) {
	query := "UPDATE users SET email = $1 WHERE id = $2"
	_, err := db.Exec(query, newEmail, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User updated successfully!")
}

func deleteUser(db *sql.DB, id int) {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully!")
}
