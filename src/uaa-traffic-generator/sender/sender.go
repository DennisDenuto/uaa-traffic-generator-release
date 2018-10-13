package sender

import "github.com/cloudfoundry-community/go-uaa"

type Sender interface {
	Send(*uaa.API)
}
