package player

import (
	"context"
	"encoding/json"
	ep "myclub/pkg/endpoints"
	"net/http"
)

func DecodeAddPlayerRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req ep.AddPlayerRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}
