package config

import (
	"encoding/json"
	"net/http"
)

func encodeResp(w http.ResponseWriter, r interface{}) error {
	return json.NewEncoder(w).Encode(r)
}

func encodeReq(w http.ResponseWriter, r interface{}) error {
	return json.NewEncoder(w).Encode(r)
}
