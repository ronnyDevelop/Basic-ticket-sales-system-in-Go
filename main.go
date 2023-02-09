package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var RemainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsletter bool
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTcketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

		if isValidName && isValidEmail && isValidTcketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			go sendTicket(userTickets, firstName, lastName, email) //go routine - concurrency

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if RemainingTickets == 0 {

				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn´t contain @ sign")
			}
			if !isValidTcketNumber {
				fmt.Println("Numbers of tickets you entered is invalid")
			}
			continue
		}

	}
}

// Functions
func greetUser() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availiable.\n", conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

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

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	RemainingTickets = RemainingTickets - userTickets

	//crate a map for user
	// var userData = make(map[string]string)

	//use struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // este es un slice
	fmt.Printf("List of bokings is %v\n", bookings)

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaninig fo %v\n", RemainingTickets, conferenceName)
}

// Simulate send ticket to email
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################")
}
