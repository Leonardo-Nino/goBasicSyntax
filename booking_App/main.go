package main

import (
	"fmt"
	"strings"
)

func main() {

	//                           -------intro app-------

	appName := "Leo first App in Go"    // 	declaration variable (Sugar, no type)
	const tickets = 50                  //  declaration constant	(no Sugar posible )
	var remainingTickets uint = tickets // 	declaration variable (tradicional, type)
	bookings := []string{}              // 	declaration slice (Sugar)

	//var bookings[]string 				//  declaration Slice(tradicional)
	//var boolings [50] string          //  declaration Array (length fixed)

	fmt.Println("Welcome to Leo's App")
	fmt.Printf("The name of the App is : %v.\n ", appName)
	fmt.Printf("Total tickets: %v, Remain ticket: %v\n  ", tickets, remainingTickets)

	//               	 		-------app logic ---------

	for { //infinity loop

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

		// Check ticket avaible

		if userTickets <= remainingTickets { //if else structure

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

			onlyFirsNames := []string{}

			for _, elementBooking := range bookings { //for each loop ("index" instead of "_" if need index)

				names := strings.Fields(elementBooking) //split function => "Leonardo Niño" ["Leonardo","Niño"]

				onlyFirsNames = append(onlyFirsNames, names[0])

			}

			fmt.Printf("All bookings: %v\n", onlyFirsNames)

			if remainingTickets == 0 {
				println("Sorry we are sold out")
				break // finnish program if no more tickest avaible
			}

		} else {

			fmt.Printf("Sorry, You can't book %v tickets beacause are only %v left \n", userTickets, remainingTickets)
			//continue start next iteration of the loop (skip the code below if true)

		}

	}

}
