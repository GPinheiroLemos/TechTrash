package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

type LogLixeira struct {
	Idlog     int     `json:"idlog"`
	Idlixeira int     `json:"idlixeira"`
	Nivel     float64 `json:"nivel"`
	Data      string  `json:"data"`
}

func GetLog(w http.ResponseWriter, r *http.Request) {

	// w.Header().Add("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()
	id, ok := query["idlixeira"]
	if !ok || len(id) == 0 {
		respError := map[string]string{"message": "missing parameter: idlixeira"}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}
	idpassado := id[0]

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT * FROM loglixeira WHERE idlixeira = %v", idpassado)
	results, err := db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	var logLixeira []LogLixeira
	for results.Next() {
		var logbanco LogLixeira
		err = results.Scan(&logbanco.Idlog, &logbanco.Idlixeira, &logbanco.Nivel, &logbanco.Data)
		if err != nil {
			respError := map[string]string{"message": "error while reading data response from database"}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
		logLixeira = append(logLixeira, logbanco)
	}

	utils.SetResponseSuccess(w, r)
	json.NewEncoder(w).Encode(logLixeira)

}
