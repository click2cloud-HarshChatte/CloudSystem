 package main

  import (
 	//"fmt"
  	"log"
 	"net/http"
     "strconv"
	  "github.com/gorilla/mux"
	  "encoding/json"
	
  )

func Numbers_Fib(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["number"]
   
	number, err := strconv.ParseInt(params, 15, 50)
	if err != nil {
	 json.NewEncoder(w).Encode(err)
	 return
	}
   
	var number0, number1, nextNumber int = 0, 1, 0
	var fib []int
   
	for i := 1; i <= int(number); i++ {
	if i == 1 {
	  fib = append(fib, number0)
	  continue
	 }
	 if i == 2 {
	  fib = append(fib, number1)
	  continue
	 }
	 nextNumber = number0 + number1
	 number0 = number1
	 number1 = nextNumber
	
	}
	fib = append(fib, nextNumber)
	json.NewEncoder(w).Encode(fib)
   }
 func main() {
	r := mux.NewRouter()
//	r.HandleFunc("/", HelloWord).Methods("GET")
	r.HandleFunc("/Numbers_Fib/{number}", Numbers_Fib)
   
	err := http.ListenAndServe(":8080", r)
	if err != nil {
	 log.Fatal("ListenAndServe: ", err)
	}
 }

// func main() {
//     r := mux.NewRouter()
// 	log.Fatal(http.ListenAndServe(":8085", r))
// }
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/hello", helloHandler).Methods("GET")
// 	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
// 	http.Handle("/", r)
// 	http.ListenAndServe(":8080", nil)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello world!")
// 	w.Header().Set("Content-Type", "application/json")
// n := mux.Vars(r)
// n, _ = strconv.Atoi(n)
// t1 := 0
// t2 := 1
// nextTerm := 0

//  fmt.Scan(&n)
//  	fmt.Print("Fibonacci Series :")
// 	for i := 1; i <= n; i++ {
//  		if i == 1 {
//  			fmt.Print(" ", t1)
//  			continue
//  		}
//  		if i == 2 {
//  			fmt.Print(" ", t2)
//  			continue
//  		}
//  		nextTerm = t1 + t2
//  		t1 = t2
//  		t2 = nextTerm

//  	}
//  	//fmt.Print(" ", nextTerm)
// }
// }

// func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Goodbye world!")
// }