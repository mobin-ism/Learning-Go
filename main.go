package main

import "fmt"

func main() {

	// Declaring fundamental variables
	conferenceName := "GO Lang Conference"
	var totalNumberOfConferenceTickets uint = 50
	var totalNumberOfConferenceTicketsSold uint = 30
	var currentBookings []string // THIS IS A SLICE. THIS IS BASICALLY AN ABSTRUCTION OF AN ARRAY. IF WE PROVIDE A LENGTH, IT WILL BE AN ARRAY. IF WE DONT PROVIDE A LENGTH, IT WILL BE A SLICE.
	

	// Greeting messages
	fmt.Printf("Welcome to our %v!\n", conferenceName)
	fmt.Printf("We have %v tickets available.\n", totalNumberOfConferenceTickets)
	fmt.Printf("We have %v tickets sold.\n", totalNumberOfConferenceTicketsSold)

	// Declaring user input variables
	var firstName string
	var lastName string
	var email string
	var numberOfTicketToPurchase uint

	fmt.Print("Enter Your Firstname: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter Your Lastname: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter Your Email: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets you want to purchase: ")
	fmt.Scan(&numberOfTicketToPurchase)

	// Greeting message to user for purchasing the tickets
	println("Thank You!", firstName, lastName, "You just bought",numberOfTicketToPurchase,"tickets. We are gonna send you a confirmation email at this email: ", email)

	// Import this booking to our current booking slice
	currentBookings = append(currentBookings, firstName +" "+ lastName)

	// Calculating the total number of remaining tickets
	totalNumberOfConferenceTicketsSold += numberOfTicketToPurchase
	totalNumberOfConferenceTickets -= totalNumberOfConferenceTicketsSold
	fmt.Printf("Now We have %v tickets available.\n", totalNumberOfConferenceTickets)
	fmt.Printf("List of users who have booked tickets: %v\n", currentBookings)
}