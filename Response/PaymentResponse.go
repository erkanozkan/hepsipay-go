 package Response
 
 type PaymentResponse struct {
	ApiKey string
 	TransactionId string
	TransactionTime string
	Signature string 
	Amount int
	Success bool
	MessageCode string
	Message string
	UserMessage string	
}