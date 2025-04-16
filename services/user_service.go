package services

import (
	"log"
	"math"

	"github.com/izacarias/lapi/domain"
)

func GetAllUsers() ([]domain.User, error) {
	users, err := domain.GetAllUsers()
	if err != nil {
		log.Printf("error getting all users: %v", err)
		return nil, err
	}
	users = enrichUsersWithLocation(users)
	users = enrichUsersWithZone(users)
	return users, nil
}

func GetUserByAddress(address string) (*domain.User, error) {
	user, err := domain.GetUserByAddress(address)
	if err != nil {
		log.Printf("error getting user by address %s: %v", address, err)
		return nil, err
	}
	user = enrichUserWithLocation(user)
	user.SetZoneId(getZoneInformation(user))
	return user, nil
}

func getZoneInformation(user *domain.User) string {
	// Get the access point ID from the user
	accessPointId := user.GetAccessPoint()
	// Get the access point by its ID
	accessPoint, err := domain.GetAccessPointById(accessPointId)
	if err != nil {
		log.Printf("error getting access point by ID %s: %v", accessPointId, err)
		return ""
	}
	// Set the zone ID in the user
	return accessPoint.GetZoneId()
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

func enrichUsersWithZone(users []domain.User) []domain.User {
	for i, user := range users {
		user.SetZoneId(getZoneInformation(&user))
		users[i] = user
	}
	return users
}

// getUserLocation retrieves the location for a user
func getUserLocation(userAddress string) *domain.Location {
	log.Printf("getting location for User %s", userAddress)
	location, err := domain.GetLocation(domain.TYPE_USER, userAddress)
	if err != nil {
		log.Printf("error getting location for user %s: %v", userAddress, err)
	}
	return location
}

func CalculateDistance(userA, userB *domain.User) (*domain.TerminalDistance, error) {
	locationA := userA.GetLocation()
	locationB := userB.GetLocation()
	distance := calculateEuclideanDistance(locationA, locationB)

	return domain.NewTerminalDistance(0, distance, 0), nil
}

func calculateEuclideanDistance(locationA, locationB *domain.Location) int {
	// Assuming locationA and locationB have Latitude and Longitude fields
	latDiff := locationA.Latitude - locationB.Latitude
	lonDiff := locationA.Longitude - locationB.Longitude
	distance := int(math.Sqrt(float64(latDiff*latDiff + lonDiff*lonDiff)))
	return distance
}
