package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var booker []string // slice
//var bookings [50]string //array
// booker := []string{} // alternative syntax for slice

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email)

			cityCase := 2
			firstNames := getFirstName()
			city := getCity(cityCase)
			fmt.Printf("These are all our boookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out, come back next year")
				fmt.Printf("See you in %v for the conference", city)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not have @ sign")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range booker { // the underscore represents a blank identifier, for variables not used. used _ instead of index
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getCity(cityCase int) string {
	city := "London"
	switch cityCase {
	case 1:
		city = "London"
	case 2, 4:
		city = "New York"
	case 3:
		city = "Singapore"
	default:
		city = "Paris"
	}
	return city
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for theri name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter Number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint,
	userTickets uint,
	firstName string,
	lastName string,
	email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName
	booker = append(booker, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmtion email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}
