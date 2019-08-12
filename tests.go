package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	userPassword := "Hola"

	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		//Todo handle exception
		log.Fatal(err)
	}

	fmt.Println("Hash to store:", string(hash))

	userPassword2 := "hola"
	hasFromDatabase := []byte("$2a$10$snl5A9Ys.9nkwzVnuaKBweZJwqGhuH/kWI0yR1wjeDszSK/eRXZcG")

	if err := bcrypt.CompareHashAndPassword(hasFromDatabase, []byte(userPassword2)); err != nil {
		//Todo handle exception
		log.Fatal(err)
	}

	fmt.Println("Password was correct")
}