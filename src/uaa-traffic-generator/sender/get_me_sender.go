package sender

import (
	"github.com/cloudfoundry-community/go-uaa"
	"log"
)

type GetMeSender struct {

}

func (GetMeSender) Send(api *uaa.API) {
	me, err := api.GetMe()
	if err != nil {
		log.Print(err)
	}
	log.Printf("GetMeSender Name: %s", me.Name)
}
