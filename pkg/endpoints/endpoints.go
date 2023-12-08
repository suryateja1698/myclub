package endpoints

import (
	playerService "myclub/pkg/services/player"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Endpoints struct {
	AddPlayerEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(logger log.Logger, playerSvc playerService.PlayerService) Endpoints {
	return Endpoints{
		AddPlayerEndpoint: AddPlayerEndpoint(logger, playerSvc),
	}
}
