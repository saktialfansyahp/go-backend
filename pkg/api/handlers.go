package api

import (
	"net/http"
	"strconv"

	"github.com/saktialfansyahp/go-rest-api/pkg/utils"
)

func HelloHandlerAdapted(w http.ResponseWriter, r *http.Request) {
    userID, _ := strconv.Atoi(r.Header.Get("X-User-ID"))
    token, _ := utils.GenerateToken(userID)

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "Hello!", "token": "` + token + `"}`))
}