package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"encoding/json"

	//"net/http"
	"github.com/gorilla/mux"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func isPalindrome(str string) interface{} {
	if str == Reverse(str) {
		return true
	}
	return false
}

func Palindrome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	json.NewEncoder(w).Encode(str)
	
        var reverseString string = ""
	var length = len(str)

	for i := length - 1; i >= 0; i -- {
		reverseString = reverseString + string(str[i])
	}

	
	if strings.ToLower(str) == strings.ToLower(reverseString) {
		fmt.Fprintln(w,"TRUE");
	} else {
		fmt.Fprintln(w,"FALSE");
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Palindrome", Palindrome).Methods("GET")
	//router.HandleFunc("/Numbers_Fib", Numbers_Fib).Methods("POST")
	log.Fatal(http.ListenAndServe(":8010", router))

}
