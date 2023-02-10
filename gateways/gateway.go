package gateways

import (
	authGateway "qurilish/api_gateway/gateways/auth"
)

type Gateway struct {
	Auth authGateway.AuthGateway
	//
}

// Init initilizes gateways to services
func Init() Gateway {
	return Gateway{
		Auth: authGateway.Init(),
	}
}
