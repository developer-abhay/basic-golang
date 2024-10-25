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












