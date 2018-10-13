package config

type TrafficConfig struct {
	UaaCommands []UaaCommand
	Credentials Credentials
}

type UaaCommand struct {
	Cmd string
}

type Credentials struct {
	Target    string
	GrantType string

	ClientId     string
	ClientSecret string

	Username     string
	UserPassword string
}
