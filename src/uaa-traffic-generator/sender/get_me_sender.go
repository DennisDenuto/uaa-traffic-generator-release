package sender

import "github.com/cloudfoundry-community/go-uaa"

type GetMeSender struct {

}

func (GetMeSender) Send(target string) {
	api, err := uaa.NewWithPasswordCredentials(target, "", "", "", "", "", uaa.JSONWebToken, true)
	if err != nil {
		panic(err)
	}

	me, err := api.GetMe()
	if err != nil {
		panic(err)
	}
	println(me.Name)
}