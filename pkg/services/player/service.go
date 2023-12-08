package player

import (
	"context"
	"encoding/json"
	"myclub/pkg/models"
	"strconv"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/segmentio/kafka-go"
)

type PlayerService interface {
	AddPlayers(ctx context.Context, firstName, lastName, country, position string, age int) error
}

type playerService struct {
	logger        log.Logger
	topic         string
	brokerAddress string
}

func NewPlayerService(logger log.Logger, topic, brokerAddress string) PlayerService {
	return &playerService{
		logger:        logger,
		topic:         topic,
		brokerAddress: brokerAddress,
	}
}

func (p *playerService) AddPlayers(ctx context.Context, firstName, lastName, country, position string, age int) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{p.brokerAddress},
		Topic:   p.topic,
		Logger:  nil,
	})

	playerDetails := models.Player{
		ID:        1,
		FirstName: firstName,
		LastName:  lastName,
		Country:   country,
		Position:  position,
		Age:       age,
	}

	data, err := json.Marshal(playerDetails)
	if err != nil {
		level.Error(p.logger).Log("message", "error while converting struct to json", "error", err)
		return err
	}

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(1)),
		Value: data,
	})
	if err != nil {
		level.Error(p.logger).Log("message", "error while writing message", "error", err)
		return err
	}
	return nil
}
