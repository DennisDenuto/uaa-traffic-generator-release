package config

type TrafficConfig struct {
	UaaCommands []UaaCommand
}

type UaaCommand struct {
	Cmd string
	Credentials Credentials
}

type Credentials struct {
	GrantType string

	ClientId string
	ClientSecret string

	Username string
	UserPassword string
}