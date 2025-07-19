package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	Db *sql.DB
}

func ConfigUserStorage(path string) (*UserStorage, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		username TEXT UNIQUE,
		password TEXT
	)
`)
	if err != nil {
		log.Println("unable tp create table in db")
		return nil, err
	}
	return &UserStorage{
		Db: db,
	}, nil
}

func (userStore *UserStorage) CreateUser(username string, email string, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	stmt, err := userStore.Db.Prepare(`INSERT INTO users (email, username, password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(email, username, hashedPassword)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (userStore *UserStorage) LoginUser(username string, password string) (bool, error) {
	stmt, err := userStore.Db.Prepare(`SELECT password FROM users WHERE username=?`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var dbPassword string
	err = stmt.QueryRow(username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("user not found")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("invalid password")
	}

	return true, nil
}
