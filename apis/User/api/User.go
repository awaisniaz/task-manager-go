package api

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am login Page")
}
func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am Register Page")
}
