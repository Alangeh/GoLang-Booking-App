package main

import "fmt"

func main() {
	conferenceName := "Go Conference"

	const conferenceTickets = 50

	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available")

	fmt.Println("Get your tickets here to attend")

	var bookings [50]string //array
	var booker []string     // slice
	// booker := []string{} // alternative syntax for slice

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
	bookings[0] = firstName + " " + lastName
	booker = append(booker, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmtion email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

	fmt.Printf("These are all our boookings: %v\n", booker)
}
