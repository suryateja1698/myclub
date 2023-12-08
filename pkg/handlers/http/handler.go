package httphandlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

type errorer interface {
	Error() error
}

type ErrorRes struct {
	ErrMessage string
}

// DefaultServerOptions returns the default server options used by transport layers.
func DefaultServerOptions(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
}

func EncodeAddPlayerResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.Error() != nil {
		encodeError(ctx, e.Error(), w)
		return e.Error()
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errResp := ErrorRes{ErrMessage: err.Error()}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errResp)
}
