package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"techTrash/connection"
)

type LogLixeira struct {
	Idlog     int     `json:"idlog"`
	Idlixeira int     `json:"idlixeira"`
	Nivel     float64 `json:"nivel"`
	Data      string  `json:"data"`
}

func GetLog(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")

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
		respError := map[string]string{"message": "mysql failed to connect"}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT * FROM loglixeira WHERE idlixeira = %v", idpassado)
	results, err := db.Query(querySQL)
	if err != nil {
		respError := map[string]string{"message": "sql query failed to execute", "query": querySQL}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
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

	if err == nil {
		resp := map[string]string{"message": "success"}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		json.NewEncoder(w).Encode(logLixeira)
		return
	}
	return
}
