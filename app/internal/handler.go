package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"demo/app/constant"
	"demo/app/util"
)

type hiRequest struct {
	Name string `json:"name"`
}

type hiResponse struct {
	Message string `json:"message"`
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	req := hiRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Cannot decode hiRequest: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := util.MergeString([]string{constant.HiMessage, req.Name}, constant.Space)
	resp, err := json.Marshal(hiResponse{Message: message})
	if err != nil {
		log.Printf("Cannot marshal response json: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("Cannot write response: %s", err.Error())
	}
}
