package user

import "errors"

type Service struct {
	repo Repository
}

func (service *Service) CreateUser(userName string) (User, error) {
	if len(userName) == 0 {
		return User{}, errors.New("userName is empty")
	}
	userInput := CreateUserInput{
		Name: userName,
	}
	return service.repo.Create(userInput)
}

func (service *Service) GetUser(i int) (User, error) {
	user, err := service.repo.Get(i)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (service *Service) DeleteUser(i int) error {
	return service.repo.Delete(i)
}

func (service *Service) GetAllUsers() []User {
	return service.repo.GetAllUsers()
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
