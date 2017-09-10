package Request

type PaymentRequest struct {
	Version string
	ApiKey string
	TransactionId string
	TransactionTime int
	Signature string
	Description string
	Amount int
	Currency string
	Installment int
	Card Card 
	Customer Customer
}

type Card struct {
	CardNumber string
	CardHolderName string
	ExpireYear string
	ExpireMonth string
	SecurityCode string
}

type Customer struct {
	IpAddress string
	Name string	 
	Surname string
	Email string
	PhoneNumber string
	Code string
	Tckn string
	VatNumber string
	MemberSince int
	Birthdate int
	City string
	District string
}
 