package main

import (
	"fmt"
	"net/http"

	"github.com/developer-abhay/probo-golang/app/handlers"
)

func main (){
	http.HandleFunc("/signup", handlers.SignUpHandler)	
	// http.HandleFunc("/signin", handlers.SignInHandler)	

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil) 
}


