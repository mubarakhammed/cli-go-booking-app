package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, firstName, lastName, conferenceName, email, bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too shaort")

			}
			if !isValidEmail {
				fmt.Println("Email is not valid")

			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is not valid ")

			}
			fmt.Printf("Your input data is invalid")

		}

	}

}

func validateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) {
	panic("unimplemented")
}

// Smaller functions

func greetUsers(confName string, confTicket int, remainTicket uint) {
	fmt.Printf("Welcome to %v booking applicatopn", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTicket, remainTicket)
	fmt.Println("Get your ticket here to attend")

}

func getFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, conferenceName string, email string, bookings []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
