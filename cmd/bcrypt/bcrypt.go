// This code is for an experimental cli app that uses bcrypt to hash the password inputed and returns the hashed password.
//
// This code is not a part of the lenslocked! I added this so I could use it as a refrence for implementing bcrypt in the actual code.

package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])

	default:
		fmt.Printf("Invalid command %q\n", os.Args[1])
	}
}

func hash(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password %v \n", password)
	}
	fmt.Printf("Password Hash: %v", string(hashedPassword))
}

func compare(hash, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Password %v is invalid! \n", password)
		return
	}
	fmt.Printf("Password is correct!")
}
