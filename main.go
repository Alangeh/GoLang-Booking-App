package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"

	const conferenceTickets = 50

	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available")

	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// ask user for theri name

	fmt.Scan(&userName)

	userTickets = 20

	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
}
