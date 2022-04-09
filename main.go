package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var confrencename = " Go Confrence"
const confrenceTicket=50
var remainingTickets=50
var bookings= make([]UserData,0)

type UserData struct
{
	firstName string
	lastName string
	email string
	numberofTickets int

}
var wg =sync.WaitGroup{}

func main(){
	greetUsers()
	fmt.Println("Get your tickets here to attend")
	for{
	firstName,lastName,email,userTicket :=getUserInput()
	isValidName,isValidemail,isValidTicketNumber := helper.ValidateUserInput(firstName,lastName,email,userTicket,remainingTickets)
	if isValidName && isValidemail && isValidTicketNumber {
		bookTicket(userTicket ,firstName ,lastName ,email )
		wg.Add(1)
		go sendTicket(userTicket ,firstName ,lastName ,email )
		firstNames :=getFirstNames()
		fmt.Printf("The first name of booking : %v \n", firstNames)
		if remainingTickets == 0	{
			fmt.Println(" Our confrence is booked out, Come back next year")
					break
		}
	
	}else if userTicket == remainingTickets {
//try some thing different
	} else 	{
		
		if !isValidName{
			fmt.Printf("Name is not valide")
	
		}
		if !isValidemail{
			fmt.Printf("Email is not valide")	
		}
		if !isValidTicketNumber{
		fmt.Printf("We only have %v tickets remaining ,so you can't book %v tickets", remainingTickets,userTicket)
		}
	}
	
	}
	wg.Wait()

}
func greetUsers(){
	fmt.Printf("Welcome to %v Booking application. \n" ,  confrencename  )
	fmt.Printf ("We have total of %v Ticket and %v are still available \n" , confrenceTicket , remainingTickets)
}
func getFirstNames() []string {
	firstNames := []string {}
		
		for _ , booking:= range bookings{
			
			firstNames= append(firstNames,booking.firstName )	
	
		}
	
	return firstNames	
}

func getUserInput() (string,string,string,int){
	var firstName string
	var lastName string
	var email string
	var userTicket int
	fmt.Print("Enter First Name :")
	fmt.Scan(&firstName)
	fmt.Print("Enter Last Name :")
	fmt.Scan(&lastName)
	fmt.Print("Enter Email Address :")
	fmt.Scan(&email)
	fmt.Print("Enter Number of tickets :")
	fmt.Scan(&userTicket)
	return firstName,lastName,email,userTicket
}
func bookTicket(userTicket int,firstName string,lastName string,email string ){
	remainingTickets =remainingTickets- userTicket
	var userData = UserData{
		firstName: firstName,
		lastName:lastName ,
		email: email,
		numberofTickets:userTicket,
	}
	

		bookings = append (bookings,userData)
		fmt.Printf("List of booking is %v\n", bookings)
		fmt.Printf("Thanks you  %v %v for booking %v. You will get the confirmation email at %v\n",firstName,lastName,userTicket,email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, confrencename)

}
func sendTicket(userTickets int, firstName string, lastName string,email string){
		time.Sleep(10 * time.Second)
			var ticket= fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)	
		fmt.Printf("\n###################################")
		fmt.Printf("\nSending ticket %v to email address %v", ticket,email)
		fmt.Printf("\n###################################")
		wg.Done()
}