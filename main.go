 package main

  import (
 	"fmt"
  	"log"
 	"net/http"
// // 	//"strconv"
  	"github.com/gorilla/mux"
  )

func Numbers_Fib(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
var n int
// // 	//n, _ = strconv.Atoi(n)
t1 := 0
t2 := 1
nextTerm := 0
var i int
 fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	
	for i = 1; i <= n; i++ {
		if i == 1 {
			//fmt.Fprint(w," ", t1)
		   continue
		} 
		 if i == 2 {
			fmt.Fprint(w," ", t2)
		continue
			} 
 		nextTerm = t1 + t2
 		t1 = t2
 		t2 = nextTerm
		 }
		
fmt.Fprint(w," The Output is ", nextTerm)
 }

 func main() {
 	router := mux.NewRouter().StrictSlash(true)
	 router.HandleFunc("/Numbers_Fib", Numbers_Fib).Methods("GET")
	 router.HandleFunc("/Numbers_Fib", Numbers_Fib).Methods("POST")
 	log.Fatal(http.ListenAndServe(":8080",router))
// 	fmt.Println("serving")
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