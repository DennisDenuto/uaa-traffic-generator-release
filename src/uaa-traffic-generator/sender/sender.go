package sender

import uaa_api "github.com/cloudfoundry-community/go-uaa"

//go:generate counterfeiter . Sender
type Sender interface {
	Send( * uaa_api.API)
}
