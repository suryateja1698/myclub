package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"myclub/pkg/models"
	clubSvc "myclub/pkg/services/club"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var (
	cmdAddPlayer = &cobra.Command{
		Use:   "add",
		Short: "add players",
		Run: func(cmd *cobra.Command, args []string) {
			clubService := clubSvc.NewClubService(Logger, Topic, BrokerAddress)
			plyrProcessor := NewPlayerProcessor(Logger, Topic, BrokerAddress)
			for {
				err := plyrProcessor.PullData(clubService)
				if err != nil {
					level.Error(Logger).Log(
						"message", "error in pulling data",
						"error", err,
					)
				}
			}
		},
	}
)

type PlayerProcessor struct {
	logger        log.Logger
	topic         string
	brokerAddress string
}

func NewPlayerProcessor(logger log.Logger, topic, brokerAddress string) *PlayerProcessor {
	return &PlayerProcessor{
		logger:        logger,
		topic:         topic,
		brokerAddress: brokerAddress,
	}
}

func (p *PlayerProcessor) PullData(svc clubSvc.ClubService) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{p.brokerAddress},
		Topic:   p.topic,
		GroupID: "first",
	})

	msg, err := r.ReadMessage(context.Background())
	if err != nil {
		level.Error(p.logger).Log(
			"message", "err in reading message",
			"error", err,
		)
	}
	var playerInfo models.Player
	err = json.Unmarshal(msg.Value, &playerInfo)
	if err != nil {
		level.Error(p.logger).Log(
			"message", "error in unmarshalling message",
			"error", err,
		)
	}
	fmt.Println("message received:", playerInfo)
	err = svc.AddPlayers(context.Background(), playerInfo)
	if err != nil {
		level.Error(p.logger).Log(
			"message", "error in adding players",
			"error", err,
		)
	}
	return nil
}
