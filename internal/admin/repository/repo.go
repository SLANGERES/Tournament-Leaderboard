package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DbConnection struct {
	Db *sql.DB
}

func ConfigAdminDB(path string) (*DbConnection, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email string, 
			username string,
			password string
		)`)

	if err != nil {
		return nil, err
	}

	return &DbConnection{Db: db}, nil
}

func (sqlDB *DbConnection) CreateAdmin(id int, email string, username string, password string) (int64, error) {
	return 0, nil
}
func (sqlDb *DbConnection) MakeUser()int64{
	return 0
}