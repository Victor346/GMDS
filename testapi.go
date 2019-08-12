package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"testApi/Middleware"
	"testApi/controllers"
)

func main() {
	router := mux.NewRouter()

	router.Use(Middleware.Authentication)

	port :=os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Listening in port ", port)

	router.HandleFunc("/", controllers.WelcomeController)
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/material/{ID}", controllers.GetOneMaterial).Methods("GET")
	router.HandleFunc("/materiales/{LowLimit}/{UpLimit}", controllers.GetMultipleMaterials).Methods("GET")

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}



}
