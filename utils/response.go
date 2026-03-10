package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, output string, err error) {

	resp := Response{
		Output: output,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
