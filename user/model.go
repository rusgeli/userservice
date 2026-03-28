package user

type User struct {
	ID   int
	Name string
}

type CreateUserInput struct {
	Name string
}
