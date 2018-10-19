package config

import "encoding/json"

const GetMeCmd = "GetMe"
const ListAllUsersCmd = "ListAllUsers"

type TrafficConfig struct {
	UaaCommands []UaaCommand
	Credentials Credentials
}

type UaaCommand struct {
	Cmd  string
	Loop int
}

type Credentials struct {
	Target    string
	GrantType string

	ClientId     string
	ClientSecret string

	Username     string
	UserPassword string
}

func NewConfig(jsonContents []byte) TrafficConfig {
	var trafficConfig TrafficConfig
	err := json.Unmarshal(jsonContents, &trafficConfig)
	if err != nil {
		panic(err)
	}

	return trafficConfig
}
