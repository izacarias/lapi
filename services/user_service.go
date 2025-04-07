package services

import "github.com/izacarias/lapi/domain"

func GetAllUsers() ([]domain.User, error) {
	return domain.GetAllUsers()
}

func GetUserByAddress(address string) (*domain.User, error) {
	return domain.GetUserByAddress(address)
}
