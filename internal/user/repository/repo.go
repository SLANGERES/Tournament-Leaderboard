package repository

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	Db *sql.DB
}

func ConfigUserStorage(dbpath string) (*UserStorage, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal("unable to connect with the db ")
		return nil, nil
	}
	_, err = db.Exec(`		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username UNIQUE TEXT,
			email UNIQUE TEXT, 
			password string
		)`)
	if err != nil {
		log.Println("unable tp create table in db")
	}
	return &UserStorage{
		Db: db,
	}, nil
}

func (userStore *UserStorage) CreateAdmin(username string, email string, password string) (int64, error) {
	//hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	stmt, err := userStore.Db.Prepare(`INSERT INTO admin (email, username, password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(username, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (userStore *UserStorage) LoginUser(username string, password string) (bool, error) {

	stmt, err := userStore.Db.Prepare(`SELECT password WHERE username=?`)
	if err != nil {

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
