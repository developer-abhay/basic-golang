package main

import (
	"fmt"
	"net/http"

	"github.com/developer-abhay/probo-golang/app/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main (){
	log.SetReportCaller(true)
	r := chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r) 
}


