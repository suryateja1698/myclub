package main

import (
	"net/http"
	"os"

	ep "myclub/pkg/endpoints"
	httpHandlers "myclub/pkg/handlers/http"

	playerService "myclub/pkg/services/player"
	playerDecorder "myclub/pkg/transport/http/player"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
)

var (
	Logger        log.Logger
	Topic         string
	BrokerAddress string
)

func Start() {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	Logger = logger
	Topic = "real-madrid"
	BrokerAddress = "localhost:9093"

	plyrSvc := playerService.NewPlayerService(Logger, Topic, BrokerAddress)

	router := mux.NewRouter()
	options := httpHandlers.DefaultServerOptions(Logger)

	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Methods("POST").Path("/players").Handler(httptransport.NewServer(
		ep.MakeServerEndpoints(Logger, plyrSvc).AddPlayerEndpoint,
		playerDecorder.DecodeAddPlayerRequest,
		httpHandlers.EncodeAddPlayerResponse,
		options...,
	))
	level.Info(Logger).Log("message", "listening on 8080")
	http.ListenAndServe(":8080", router)
}
