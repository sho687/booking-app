package main

import "strings"

func ValidateUser(firstName string, secondName string, email string, userTickets uint, remaining_tickets uint) (bool, bool, bool) { //  function checks if user data is valid

	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTickets > 0 && userTickets <= remaining_tickets
	return isValidName, isvalidEmail, isvalidTicketNumber
}
