package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "admin123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password:", password)
	fmt.Println("Hash:", string(hashedPassword))

	// Verificar que funciona
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		log.Fatal("Verification failed")
	}
	fmt.Println("Verification: OK")
}
