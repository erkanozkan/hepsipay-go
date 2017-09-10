package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"io/ioutil" 
	"./Request"  
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HepsiPay Go Client!")
}

func Payment(w http.ResponseWriter, r *http.Request) {
	var request Request.PaymentRequest
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		
	if err != nil {
		respondJSON(w, 422, err)		
	}
  
	if err := json.Unmarshal(body, &request); err != nil {
		respondJSON(w, 422, err)
	} 

	b,err := json.Marshal(request)

	response := HepsiPaymentApi(b) 
	 
	respondJSON(w, http.StatusOK, response)
}
