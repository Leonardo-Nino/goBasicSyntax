package helper

import "strings"

func ValidImputs(userName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidNumberTicket := userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidNumberTicket
}
