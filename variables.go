package main

//Importing methods of library fmt
import (
	"fmt"
)

func main() {
	// 1 - Calling function = helloWorld()
	helloWorld()
	// 2 - Calling function = printNumber()
	printNumber()
	// 3 - calling function = sumTwoNumbers()
	sumTwoNumbers()
	// 4 - calling function = fourNotesAverage()
	fourNotesAverage()

}

// =========================================== section of functinos for calling in main function ===========================================

func helloWorld() {
	fmt.Println("Hello World, Francisco!!")
}

func printNumber() {
	var number int
	fmt.Println("Enter with your number: ")
	fmt.Scan(&number)

	fmt.Printf("Your number is %v\n", number)
}

func sumTwoNumbers() {
	var numberOne int
	var numberTwo int

	fmt.Println("Enter with a number one: ")
	fmt.Scan(&numberOne)
	fmt.Println("Enter with a number two: ")
	fmt.Scan(&numberTwo)

	var sumTwoNumbers = numberOne + numberTwo

	fmt.Printf("The sum of numbers is: %v\n", sumTwoNumbers)
}

func fourNotesAverage() {
	var noteOne uint
	var noteTwo uint
	var noteThree uint
	var noteFour uint
	var averageNotes uint

	//Input data user
	fmt.Println("Note one: ")
	fmt.Scan(&noteOne)

	fmt.Println("Note two: ")
	fmt.Scan(&noteTwo)

	fmt.Println("Note three: ")
	fmt.Scan(&noteThree)

	fmt.Println("Note four: ")
	fmt.Scan(&noteFour)

	averageNotes = (noteOne + noteTwo + noteThree + noteFour) / 4

	fmt.Printf("The you average notes is: %v", averageNotes)

}
