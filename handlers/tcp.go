package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/prashantjay235/willitconnect/utils"
)

type TCPRequest struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func TCPCheck(w http.ResponseWriter, r *http.Request) {

	var req TCPRequest
	json.NewDecoder(r.Body).Decode(&req)

	cmd := exec.Command("nc", "-vz", req.Host, req.Port)

	out, err := cmd.CombinedOutput()

	utils.WriteJSON(w, string(out), err)

}
