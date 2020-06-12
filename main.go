package main

import (
	"fmt"
	"net/http"
	//"log"
	//"net/http"
)

func New(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (New(n-1) + New(n-2))
	}
}

func Compare(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var n, i, j, k int
	j = 0
	fmt.Print("Enter the numbers")
	fmt.Scanln(&n)
	fmt.Print("First Series to be Printed")
	for i = 1; i <= n; i++ {
		fmt.Print(New(j), ",")
		j++
	}
	fmt.Println()

	fmt.Println("Second Series to be Printed")
	for k = 1; k <= 10; k++ {
		fmt.Print(k, ",")
	}

}
