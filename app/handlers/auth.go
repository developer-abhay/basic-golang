package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/developer-abhay/probo-golang/app/db"
	"github.com/developer-abhay/probo-golang/app/models"
)

// Signup Route
func SignUpHandler (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Invalid Request Method")
		return
	}

	// Decode r.body and store in new user
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// Check format of the r.body 
	if newUser.Name == ""  {
		response := map[string]string{"error":"Name required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if newUser.Email == ""  {
		response := map[string]string{"error":"Email required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if newUser.Password == ""  {
		response := map[string]string{"error":"Password required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if user already exists
	if _,exists := db.Users[newUser.Email]; exists {
		response := map[string]string{"error":"User Already exists"}
		
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Create a new user
	db.Users[newUser.Email] = newUser
	response := map[string]string{"message":"User Created successfully"}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


// Singin Route
func SignInHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	// Only allow post method requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Invalid Request Method")
		return
	}
	
	// Store r.body into varible user 
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	
	// Check for valid email and password
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error":"Please enter email and password"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Check if user exist with the given email
	userDetails,userExists := db.Users[user.Email];
	
	if  !userExists {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error":"User does not exist"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Verify email and password
	if userDetails.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		response := map[string]string{"error":"Email and password does not match"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	response := map[string]string{"message":"Login successfully"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
