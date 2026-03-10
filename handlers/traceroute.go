package handlers

import (
	"net/http"
	"os/exec"

	"github.com/prashantjay235/willitconnect/utils"
)

func TraceRoute(w http.ResponseWriter, r *http.Request) {

	host := r.URL.Query().Get("host")

	cmd := exec.Command("traceroute", host)

	out, err := cmd.CombinedOutput()

	utils.WriteJSON(w, string(out), err)

}
