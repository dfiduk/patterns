package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	fmt.Println(currentUser.Username == "root")
}
