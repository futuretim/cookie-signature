package main

import (
	"fmt"
	cookie "gogs.base127.com/cookie-signature"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		fmt.Println("Please supply a message and a secret")
		os.Exit(1)
	}

	secret := argsWithoutProg[0]
	message := argsWithoutProg[1]

	replacer := strings.NewReplacer("\\", "", "=", "", "$", "",)

	signedCookie := cookie.Signcookie(message, secret, replacer)
	complete := message + "." + signedCookie
	fmt.Println("Signed cookie: " + complete)

	if cookie.Unsigncookie(complete, secret, replacer) {
		fmt.Println("It's good!")
	} else {
		fmt.Println("It's NOT so good!")
	}
}