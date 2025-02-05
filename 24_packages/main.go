package main

import (
	"fmt"
	// "os/user"

	"github.com/tushar07-debug/podcast/auth"
	"github.com/tushar07-debug/podcast/user"
)

func main() {
	auth.Login("tushar", "secretcd")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{Email: "tushar@gmail.com", Name: "Tushar SIngh"}
	fmt.Println("user", user.Email, user.Name)
}
