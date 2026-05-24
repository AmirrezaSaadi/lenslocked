// This code is for an experimental cli app that uses bcrypt to hash the password inputed and returns the hashed password.
//
// This code is not a part of the lenslocked! I added this so I could use it as a refrence for implementing bcrypt in the actual code.

package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash := "password"
		fmt.Printf("hash = %q\n", hash)
	case "compare":
		fmt.Printf("comparison of hash and password: True")

	default:
		fmt.Printf("Invalid command %q\n", os.Args[1])
	}
}

func hash(password string) {

}

func compare(password, hash string) {

}
