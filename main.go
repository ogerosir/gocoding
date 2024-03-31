package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName = "BioInformatics Seasson 1"

var conferenceTickets uint = 50
var remainingTickets uint = 50

// var bookings []string // an empty slice with list of strings
var bookings = make([]UserData, 0) // an empty slice with list of maps

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}
var wg = sync.WaitGroup{}
func main() {
	// var bookings =[] string{}
	// bookings :=[] string{}
	greetUsers()

	// for len(booking)<50{ run code here...}
	// for {
	// getUserInput()
	firstName, lastName, email, userTickets := getUserInput()
	// user input validation
	isValidName, isValidEmail, isValidTicketNumber := UserInputValidation(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		// bookTicket()
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		// FirstNames(bookings)
		firstNames := FirstNames()
		fmt.Printf("The first name is %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Conference fuly booked, try next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("Sorry, first name or last name is too short.")
		}
		if !isValidEmail {
			fmt.Println("Sorry, the email must have the @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Sorry, the number of tickets you entered is invalid.")
		}
	}

	// }

	wg.Wait()
}

func greetUsers() {
	fmt.Println("Helo Friend, Welcome to Ticket Booking App")
	fmt.Printf("Get your tickets for %v here\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available.\n", conferenceTickets, remainingTickets)
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []UserData {
	remainingTickets = remainingTickets - userTickets
	//create object of UserData
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is %v ", bookings)
	fmt.Printf("We have %v booking\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v you will get confirmation massage at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("The remaining tickets now are %v\n", remainingTickets)
	return bookings
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v \n to emeil address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}
