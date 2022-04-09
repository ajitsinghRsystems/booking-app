package helper
import "strings"

func ValidateUserInput(firstName string,lastName string,email string,userTicket int, remainingTickets int) (bool,bool,bool) {

	isValidName:= len(firstName)>2 && len(lastName)>2
	isValidemail:= strings.Contains(email,"@")
	isValidTicketNumber:= userTicket>0 && userTicket< remainingTickets
	return isValidName,isValidemail,isValidTicketNumber
}