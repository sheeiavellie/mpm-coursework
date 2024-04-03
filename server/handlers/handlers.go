package handlers

import (
	"log"
	"net/http"

	"github.com/sheeiavellie/mpm-coursework/server/data"
	"github.com/sheeiavellie/mpm-coursework/server/util"
)

func HandleProcess(w http.ResponseWriter, r *http.Request) {
	var request data.ProcessRequest
	err := util.ReadJSON(r, &request)
	if err != nil {
		log.Printf("Invalid request: %s", err)
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, "Process recived.")
	if err != nil {
		log.Printf("Can't write json: %s", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	log.Printf("Process recived, timestamp: %s\n", request.Timestamp)
	log.Printf("Data:\n")
	for _, sensor := range request.Data {
		log.Printf("%s", sensor.GetSensorData())
	}

}
