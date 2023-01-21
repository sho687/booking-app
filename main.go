package main

//importing package

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "GO conference" //conference name
const conference_tickets = 50        //total tickets
var remaining_tickets uint = 50
var bookings = make([]userdata, 0)

type userdata struct {
	firstName   string
	secondName  string
	email       string
	userTickets uint
}

var waiting = sync.WaitGroup{}

func main() { //main function starts

	greetUsers() //greeting users
	for {

		firstName, secondName, email, userTickets := userInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUser(firstName, secondName, email, userTickets, remaining_tickets)

		if isValidName && isValidEmail && isValidTicketNumber { //runs the program if user data is valid

			bookTicket(userTickets, firstName, secondName, email) //function call to book ticket
			waiting.Add(1)
			go sendTicket(userTickets, firstName, secondName, email) //function call to send ticket
			getFirstNames()
			//firstNames := getFirstNames()s
			var printfname = getFirstNames()
			fmt.Printf("the first names of bookings are : %v  \n", printfname)
			//fmt.Print(firstNames, "\n")

			if remaining_tickets == 0 { //Ends the app if all tickets are sold

				fmt.Println("Our conference is booked out.Come back next year!")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("firstname or second name you entered is too short ")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The  Ticket number you entered is invalid")
			}

			//fmt.Println("Your user input data is invalid, try again") //displays this message if user input data is invalid
			continue
		}

	}
	waiting.Wait()
}

func greetUsers() {
	//Start of the Booking application
	fmt.Printf("Welcome to %v booking application \n ", conferenceName)
	fmt.Printf("we have total of %v tickets and still  %v are available \n ", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here !")
}

func getFirstNames() []string { //function to get first names
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	//fmt.Printf("First names of bookings are : %v\n", firstNames)
	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var userTickets uint
	// User's data (input)
	fmt.Println("Enter your first  name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your secoond name:")
	fmt.Scan(&secondName)

	fmt.Println("Enter your email Id")
	fmt.Scan(&email)

	fmt.Println("Number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, secondName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, secondName string, email string) {
	remaining_tickets = remaining_tickets - userTickets

	//creating a map
	var userData = userdata{
		firstName:   firstName,
		secondName:  secondName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, secondName, userTickets, email)
	fmt.Printf("Only %v tickets are remaining for %v \n ", remaining_tickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, secondName string, email string) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for  %v %v\n ", userTickets, firstName, secondName)
	fmt.Println("#####################")
	fmt.Printf("sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("#####################")
	waiting.Done()
}
