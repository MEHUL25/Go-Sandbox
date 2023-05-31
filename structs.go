package main

import (
	"fmt"
)

type mail struct {
	message  string
	sender   user
	receiver user
}

type user struct {
	name string
}

func can_send(test_mail mail) bool {
	if test_mail.sender.name == "" {
		return false
	}
	if test_mail.receiver.name == "" {
		return false
	}
	return true
}

func main() {
	test_mail := mail{}

	test_mail.sender.name = "Mehul"
	//test_mail.receiver.name = "Prakash"

	if can_send(test_mail) {
		fmt.Println("Mail is valid")
	} else {
		fmt.Println("Mail is not valid")
	}

}
