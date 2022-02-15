package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/gorilla/mux"

	"com.damico958.devops/calculator"
)

const byte = 64
const port = ":8080"

func CalculationHandler(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	firstNumber, errorFirstNumber := strconv.ParseFloat(parameters["a"], byte)
	secondNumber, errorSecondNumberNumber := strconv.ParseFloat(parameters["b"], byte)
	chosenOperation := parameters["op"]

	if errorFirstNumber != nil {
		fmt.Fprintf(w, "Fail! Could not convert this first number. Try again!")
	} else if errorSecondNumberNumber != nil {
		fmt.Fprintf(w, "Fail! Could not convert this second number. Try again!")
	} else {
		json.NewEncoder(w).Encode(calculator.Calculate(firstNumber, secondNumber, chosenOperation))
		log.Println("Operation correctly performed!")
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(calculator.History)
	log.Println("History correctly checked!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calc/{op}/{a}/{b}", CalculationHandler).Methods("GET")
	router.HandleFunc("/calc/history", HistoryHandler).Methods("GET")
	log.Println("Listening on " + port + " port")
	log.Fatal(http.ListenAndServe(port, router))
}
