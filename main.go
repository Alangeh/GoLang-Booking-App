package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"

	const conferenceTickets = 50

	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available")

	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmtion email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
