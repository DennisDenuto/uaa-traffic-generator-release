package sender

import (
	"github.com/cloudfoundry-community/go-uaa"
)

type GetMeSender struct {

}

func (GetMeSender) Send(api *uaa.API) {
	me, err := api.GetMe()
	if err != nil {
		panic(err)
	}
	println(me.Name)
}