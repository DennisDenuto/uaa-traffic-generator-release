package sender

import "cloudfoundry-community/go-uaa/go-uaa"

type Sender interface {
	Send(target string)
}

func Send(target string) {
	api, err := uaa.NewWithPasswordCredentials(target, "", "", "", "", "", uaa.JSONWebToken)
	if err != nil {
		panic(err)
	}

	me, err := api.GetMe()
	if err != nil {
		panic(err)
	}
	println(me.Name)
}