package service

type UserStorage interface {
	CreateUser(username string, email string, passowrd string) (int64, error)
	LoginUser(username string, password string) (bool, error)
}
