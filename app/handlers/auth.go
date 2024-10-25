package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/developer-abhay/probo-golang/app/db"
	"github.com/developer-abhay/probo-golang/app/models"
)

// Signup ROute
func SignUpHandler (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Invalid Request Method")
		return
	}

	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	if _,exists := db.Users[newUser.Email]; exists {
		fmt.Print("User Already exists")
		return
	}
	
	db.Users[newUser.Email] = newUser
	fmt.Print("User Created successfully")
}


















// Signup Route
// func SignUpHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Check if request is POST
// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode("Invalid request method")
// 		return
// 	}
	
// 	var newUser models.User
// 	json.NewDecoder(r.Body).Decode(&newUser)
	
// 	if newUser.Email == "" || newUser.Name == "" || newUser.Password == ""   {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := map[string]string{"error": "Invalid Inputs"}
		
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Check if user already exists
// 	if _,exist := db.Users[newUser.Email] ; exist {
// 		w.WriteHeader(http.StatusConflict)
// 		response := map[string]string{"error": "User already exists"}
		
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
	
// 	// Create a new user
// 	db.Users[newUser.Email] = newUser
	
// 	w.WriteHeader(http.StatusOK)
// 	response := map[string]string{"message":"User created Successfully"}
	
// 	json.NewEncoder(w).Encode(response)
// }

// SignIn Route
// func SignInHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w,"Invalid request method" , http.StatusMethodNotAllowed)
// 		return
// 	}

// 	path := r.URL.Path

// 	userId := strings.TrimPrefix(path,"/signin/")
	

// 	if strings.Contains(userId, "/") {
// 		fmt.Println("Page not found")
// 		return
// 	}

// 	fmt.Println(userId)
// 	// fmt.Println(url.Get("name"))
// 	// fmt.Println(url.Get("age"))
// }