package service

type AdminService struct {
	store Storage
}

func NewUserService(store Storage) *AdminService {
	return &AdminService{store: store}
}

func (s *AdminService) SignupAdmin(username, email, password string) (int64, error) {
	return s.store.CreateAdmin(username, email, password)
}

func (s *AdminService) LoginAdmin(username, password string) (int64, error) {
	return s.store.LoginAdmin(username, password)
}
