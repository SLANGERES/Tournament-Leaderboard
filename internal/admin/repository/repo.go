package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type DbConnection struct {
	Db *sql.DB
}

func ConfigAdminDB(path string) (*DbConnection, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS admin (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		username TEXT UNIQUE,
		password TEXT
	)
`)

	if err != nil {
		return nil, err
	}

	return &DbConnection{Db: db}, nil
}

func (sqlDB *DbConnection) CreateAdmin(email, username, password string) (int64, error) {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	stmt, err := sqlDB.Db.Prepare(`INSERT INTO admin (email, username, password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(email, username, string(hashedPassword))
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (sqlDB *DbConnection) LoginAdmin(username, password string) error {
	stmt, err := sqlDB.Db.Prepare(`SELECT password FROM admin WHERE username = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var dbPassword string
	err = stmt.QueryRow(username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))

	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}
