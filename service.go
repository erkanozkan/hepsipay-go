package main

import ( 
	"bytes" 
	"log"
	"net/http"
	"fmt" 
	"io/ioutil" 
	"encoding/json"	 
	"./Response" 
)

func HepsiPaymentApi(paymentRequest []byte) Response.PaymentResponse{
	url := ReadConfigStr().PaymentUrl
	 
    fmt.Println("URL:>", url)
 
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(paymentRequest))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
 
	response := Response.PaymentResponse{}
	
	body, _ := ioutil.ReadAll(resp.Body)
	
	jsonErr := json.Unmarshal(body, &response)

	if jsonErr != nil {
		log.Fatal(jsonErr)		
	}  
    return response 
}