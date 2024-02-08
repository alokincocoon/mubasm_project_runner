package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"regexp"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

func validateWithLuhnAlogo(number string) bool {
	isSecondDigit := false
	sumOfNumbers := 0
	for i := 0; i < len(number); i++ {
		num := int(number[i] - '0')

		if isSecondDigit {
			num *= 2
			if num > 9 {
				num = (num % 10) + 1
			}
			
		}
		isSecondDigit = !isSecondDigit
		sumOfNumbers += num
	}
	return (sumOfNumbers % 10 == 0)

}


func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func validateCreditCard(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	number := vars["number"]
	regex, _ := regexp.Compile("^[0-9]+$")
	fmt.Println(regex.MatchString(number))
	if !regex.MatchString(number) {
		w.WriteHeader(http.StatusBadRequest)
		resp := make(map[string]string)
		resp["message"] = "Bad Request: Please provide a card number"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}
	isValid := validateWithLuhnAlogo(number)
	if isValid {
		fmt.Fprintf(w, "Provided credit card is valid according to Luhn Algorithm")
	} else {
		fmt.Fprintf(w, "Provided credit card is not valid according to Luhn Algorithm")
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/validate/{number}", validateCreditCard)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), myRouter))
}

func main() {
	godotenv.Load()
	handleRequests()
}
