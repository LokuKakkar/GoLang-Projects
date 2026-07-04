package main

import (
	"fmt"
	"log"
)

// User is a struct that represents a user in the database
type User struct {
	ID        int
	FirstName string
}

type MockDataStore struct {
	Users map[int]User
}

// GetUserByID is a method that returns a user by their ID
func (md MockDataStore) GetUserByID(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

// CreateUser is a method that creates a new user
func (md MockDataStore) CreateUser(user User) error {
	_, ok := md.Users[user.ID]
	if ok {
		return fmt.Errorf("user %v already exists", user.ID)
	}
	md.Users[user.ID] = user
	return nil
}

// DataStore is an interface that defines the methods for a data store
type DataStore interface {
	GetUserByID(id int) (User, error)
	CreateUser(user User) error
}

func main4() {
	fmt.Println("Hello World 4")
	log.Println("This is a log message")

}
