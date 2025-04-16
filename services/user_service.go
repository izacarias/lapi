package services

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

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
	timeStamp := calculateTimestamp(locationA, locationB)

	return domain.NewTerminalDistance(0, distance, timeStamp), nil
}

func CalculateDistanceLatLong(user *domain.User, latitude, longitude string) (*domain.TerminalDistance, error) {
	// Convert latitude and longitude to float32
	lat, err := strconv.ParseFloat(latitude, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid latitude: %v", err)
	}
	lon, err := strconv.ParseFloat(longitude, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid longitude: %v", err)
	}
	locationA := user.GetLocation()
	locationB := &domain.Location{
		Latitude:  float32(lat),
		Longitude: float32(lon),
		// Since the altitude is not provided, we can set it to the user's altitude
		Altitude: locationA.Altitude,
	}
	distance := calculateEuclideanDistance(locationA, locationB)
	timeStamp := calculateTimestamp(locationA, locationB)
	return domain.NewTerminalDistance(0, distance, timeStamp), nil
}

func calculateEuclideanDistance(locationA, locationB *domain.Location) int {
	// Assuming locationA and locationB have Latitude and Longitude fields
	latDiff := locationA.Latitude - locationB.Latitude
	lonDiff := locationA.Longitude - locationB.Longitude
	altDiff := locationA.Altitude - locationB.Altitude
	// Calculate the Euclidean distance
	distance := int(math.Sqrt(float64(latDiff*latDiff + lonDiff*lonDiff + altDiff*altDiff)))
	return distance
}

func calculateTimestamp(locationA, locationB *domain.Location) time.Time {
	if locationA.GetTimestamp().Before(locationB.GetTimestamp()) {
		return locationA.GetTimestamp()
	}
	return locationB.GetTimestamp()
}
