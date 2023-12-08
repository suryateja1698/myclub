package endpoints

import (
	"context"
	"myclub/pkg/errors"
	playerService "myclub/pkg/services/player"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type AddPlayerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Position  string `json:"position"`
	Age       int    `json:"age"`
}

type AddPlayerResponse struct {
	Produced bool `json:"produced"`
}

func AddPlayerEndpoint(logger log.Logger, playerSvc playerService.PlayerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddPlayerRequest)
		if err := validateRequest(logger, req); err != nil {
			level.Error(logger).Log(
				"action", "pkg.endpoints.AddPlayerEndpoint",
				"message", "error while validating request body",
				"request_body", req,
				"error", err,
			)
			return nil, err
		}

		err = playerSvc.AddPlayers(ctx, req.FirstName, req.LastName, req.Country, req.Position, req.Age)
		if err != nil {
			return nil, err
		}

		return AddPlayerResponse{
			Produced: true,
		}, nil
	}
}

func validateRequest(logger log.Logger, req AddPlayerRequest) error {
	if req.Age == 0 {
		level.Error(logger).Log("message", "empty age")
		return errors.ErrInvalidAge
	}

	if req.Position == "" {
		level.Error(logger).Log("message", "position is empty", "name", req.FirstName+" "+req.LastName)
		return errors.ErrEmptyPosition
	}
	return nil
}
