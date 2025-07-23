package service

type UserService struct {
	store UserStorage // this is the interface
}

func NewUserService(store UserStorage) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Signup(username, email, password string) (int64, error) {
	return s.store.CreateUser(username, email, password)
}

func (s *UserService) Login(username, password string) (int64, error) {
	return s.store.LoginUser(username, password)
}
