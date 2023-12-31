package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var booker = make([]map[string]string, 0) // slice of maps
var booker = make([]UserData, 0) // slice of struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//var bookings [50]string //array
// booker := []string{} // alternative syntax for slice

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email)

			wg.Add(1) // 1 is number of threads
			go sendTicket(userTickets, firstName, lastName, email) // go keyword implements asynchronous programming
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range booker { // the underscore represents a blank identifier, for variables not used. used _ instead of index
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
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

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	booker = append(booker, userData)
	fmt.Printf("List of bookings is %v\n", booker)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmtion email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket :\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}
