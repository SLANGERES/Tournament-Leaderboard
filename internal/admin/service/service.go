package service

type Storage interface{
	CreateAdmin(id string, email string, username string, password string)(int64,error)
	LoginUser(username string, password string)
	MakeUser()int64

}