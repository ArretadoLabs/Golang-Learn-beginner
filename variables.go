package main

//Importing methods of library fmt
import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Print("Printing type of data store in variables\n")

	//Using %T for discover of type data store in variable
	fmt.Printf("ConferenceTickets is %T, remainingTickets %T and conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("============= Printing values with on format interpolation =============")

	fmt.Printf("Welcome to %v booking application\n", conferenceName) //Concept of interpolation and explict new line "\n"
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("================= manipulation variables =================")

	//declaring variables
	var userName string
	var userTickets int

	userName = "Francisco"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	fmt.Println("================= manipulation of variables with pointers ==================")
	// Section of create variables
	var firstName string
	var lastName string
	var email string
	var ticketsUser uint

	// Using function "Scan" for read input data user
	fmt.Println("What your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("What your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("What your email contact: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&ticketsUser)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, ticketsUser, email)
}
