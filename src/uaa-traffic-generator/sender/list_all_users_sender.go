package sender

import "github.com/cloudfoundry-community/go-uaa"

type ListAllUsersSender struct {}

func (ListAllUsersSender) Send(*uaa.API) {
	panic("implement me")
}

