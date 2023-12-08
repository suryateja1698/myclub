package club

import (
	"context"
	"myclub/pkg/models"

	"github.com/go-kit/log"
)

var (
	players = make(map[int]models.Player)
)

type ClubService interface {
	AddPlayers(ctx context.Context, playerInfo models.Player) error
}

type clubService struct {
	logger        log.Logger
	topic         string
	brokerAddress string
}

func NewClubService(logger log.Logger, topic, brokerAddress string) ClubService {
	return &clubService{
		logger:        logger,
		topic:         topic,
		brokerAddress: brokerAddress,
	}
}

func (c *clubService) AddPlayers(ctx context.Context, playerInfo models.Player) error {
	players[playerInfo.ID] = playerInfo
	return nil
}
