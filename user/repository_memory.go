package user

import (
	"errors"
)

type InMemoryRepo struct {
	data   map[int]User
	nextID int
}

func (repo *InMemoryRepo) Create(userInput CreateUserInput) (User, error) {
	user := User{
		ID:   repo.nextID,
		Name: userInput.Name,
	}
	repo.nextID++
	repo.data[user.ID] = user
	return user, nil
}

func (repo *InMemoryRepo) Get(id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("user index out of range")
	}
	if _, ok := repo.data[id]; ok {
		return repo.data[id], nil
	}
	return User{}, errors.New("user does not exist")
}

func (repo *InMemoryRepo) Delete(id int) error {
	if id <= 0 {
		return errors.New("user index out of range")
	}
	if _, ok := repo.data[id]; ok {
		delete(repo.data, id)
		return nil
	}
	return errors.New("user does not exist")
}

func (repo *InMemoryRepo) GetAllUsers() []User {
	var result = make([]User, 0)
	for _, value := range repo.data {
		result = append(result, value)
	}
	return result
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data:   make(map[int]User),
		nextID: 1,
	}
}
