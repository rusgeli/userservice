package user

type Repository interface {
	Create(userInput CreateUserInput) (User, error)
	Get(id int) (User, error)
	Delete(id int) error
	GetAllUsers() []User
}
