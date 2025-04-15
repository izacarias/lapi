package services

import (
	"log"

	"github.com/izacarias/lapi/domain"
)

func GetAllUsers() ([]domain.User, error) {
	users, err := domain.GetAllUsers()
	if err != nil {
		log.Printf("error getting all users: %v", err)
		return nil, err
	}
	return enrichUsersWithLocation(users), nil
}

func GetUserByAddress(address string) (*domain.User, error) {
	user, err := domain.GetUserByAddress(address)
	if err != nil {
		log.Printf("error getting user by address %s: %v", address, err)
		return nil, err
	}
	user = enrichUserWithLocation(user)
	return user, nil
}

func enrichUserWithLocation(user *domain.User) *domain.User {
	user.SetLocation(getUserLocation(user.GetAddress()))
	return user
}

func enrichUsersWithLocation(users []domain.User) []domain.User {
	for i, user := range users {
		user.SetLocation(getUserLocation(user.GetAddress()))
		users[i] = user
	}
	return users
}

// getAPLocation retrieves the location for an access point
func getUserLocation(userAddress string) *domain.Location {
	log.Printf("getting location for User %s", userAddress)
	location, err := domain.GetLocation(domain.TYPE_USER, userAddress)
	if err != nil {
		log.Printf("error getting location for user %s: %v", userAddress, err)
	}
	return location
}
