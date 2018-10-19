package sender

import (
	"github.com/cloudfoundry-community/go-uaa"
	"log"
)

type ListAllUsersSender struct {}

func (ListAllUsersSender) Send(api *uaa.API) {
	users, err := api.ListAllUsers("", "", "", uaa.SortAscending)
	if err != nil {
		log.Print(err)
	}

	for _, user := range users {
		log.Printf("ListAllUsersSender returned users: %v", user)
	}
}

