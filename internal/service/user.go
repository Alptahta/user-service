package service

type UserRepository interface{}

type UserService interface{}

type User struct {
	repo UserRepository
}

func NewUserService(userRepository UserRepository) *User {
	return &User{
		repo: userRepository,
	}

}
