package main

import (
	"fmt"
	cookie "gogs.base127.com/cookie-signature"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		fmt.Println("Please supply a message and a secret")
		os.Exit(1)
	}

	secret := argsWithoutProg[0]
	message := argsWithoutProg[1]

	signedCookie := cookie.Signcookie(message, secret)
	complete := message + "." + signedCookie
	fmt.Println("Signed cookie: " + complete)

	if cookie.Unsigncookie(complete, secret) {
		fmt.Println("It's good!")
	} else {
		fmt.Println("It's NOT so good!")
	}
}