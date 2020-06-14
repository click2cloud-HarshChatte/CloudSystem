package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"encoding/json"

	
	"github.com/gorilla/mux"
)



func Palindrome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	fmt.Fprintf(w,"Enter a string: ")
	 params := mux.Vars(r)
	 name:= params["str"]
	
	json.NewEncoder(w).Encode(name)
	
        var reverseString string = ""
	var length = len(name)

	for i := length - 1; i >= 0; i -- {
		reverseString = reverseString + string(name[i])
	}

	
	if strings.ToLower(name) == strings.ToLower(reverseString) {
		fmt.Fprintln(w,"TRUE");
	} else {
		fmt.Fprintln(w,"FALSE");
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Palindrome/{str}", Palindrome)	
	log.Fatal(http.ListenAndServe(":8010", router))
}
