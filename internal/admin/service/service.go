package service

type Storage interface {
	CreateAdmin(email string, username string, password string) (int64, error)
	LoginAdmin(username string, password string) (int64, error)
}
