package main

import (
	"fmt"
	"strings"
)

// globals vars

const tickets = 50 //  declaration constant	(no Sugar posible )
var appName string = "Leo first App in Go"
var remainingTickets uint = tickets // 	declaration variable (tradicional, type)
var bookings = []string{}           // 	declaration slice (tradicional )

//alternative syntax

//appName := "Leo first App in Go"    // declaration variable (Sugar, no type)
//bookings := []string 				//   declaration Slice(tradicional)
//var boolings [50] string          //   declaration Array (length fixed)

func main() {

	//                           -------intro app-------

	greetings()

	//               	 		-------app logic ---------

	for { //infinity loop if no condition

		var userName string
		var lastname string
		var userEmail string
		var userTickets uint

		//ask user info

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&userName) // pointer!!!!!!!! Scan inputs

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastname)

		fmt.Println("Please enter your email: ")
		fmt.Scan(&userEmail)

		fmt.Println("Please number of tickets: ")
		fmt.Scan(&userTickets)

		// Check ticket avaible and data imput

		isValidName, isvalidEmail, isValidNumberTicket := validImputs(userName, lastname, userEmail, userTickets)

		// login for Bookings

		if isValidName && isvalidEmail && isValidNumberTicket { //if else structure

			// Decrease  ticket avaible

			remainingTickets -= userTickets

			// Add elements to booking Slice

			bookings = append(bookings, userName+" "+lastname) //append to Slice

			fmt.Printf("Type of Slice: %T \n", bookings) //type of
			fmt.Printf("Slice elements: %v  \n", bookings)
			fmt.Printf("Slice element at pos [0]: %v  \n", bookings[0])
			fmt.Printf("Slice length: %v  \n", len(bookings)) //length of array/slice

			fmt.Printf("Thanks %v %v. You booked %v tickets. You will receive a confirmation at %v \n", userName, lastname, userTickets, userEmail)
			fmt.Printf("Total tickets: %v, Remain ticket: %v\n", tickets, remainingTickets)

			// Print first name

			firstNameToPrint := firstName()
			fmt.Printf("All bookings: %v\n", firstNameToPrint)

			if remainingTickets == 0 {
				println("Sorry we are sold out")
				break // finnish program if no more tickest avaible
			}

		} else {
			if !isValidName {
				fmt.Printf("Sorry, your first name or last name is to short \n")
			}

			if !isValidNumberTicket {
				fmt.Printf("Sorry, You can't book %v tickets because are only %v left \n", userTickets, remainingTickets)
				//continue start next iteration of the loop (skip the code below if true)
			}

			if !isvalidEmail {
				fmt.Printf("Sorry, your Email have incorrect input\n")

			}

		}

	}

}

func greetings() {
	fmt.Println("Welcome to Leo's App")
	fmt.Printf("The name of the App is : %v.\n ", appName)
	fmt.Printf("Total tickets: %v, Remain ticket: %v\n  ", tickets, remainingTickets)
}

func firstName() []string {
	onlyFirsNames := []string{}
	for _, elementBooking := range bookings { //for each loop ("index" instead of "_" if need index)
		names := strings.Fields(elementBooking) //split function => "Leonardo Niño" ["Leonardo","Niño"]
		onlyFirsNames = append(onlyFirsNames, names[0])
	}
	return onlyFirsNames
}

func validImputs(userName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(userEmail, "@")
	isValidNumberTicket := userTickets <= remainingTickets
	return isValidName, isvalidEmail, isValidNumberTicket
}
