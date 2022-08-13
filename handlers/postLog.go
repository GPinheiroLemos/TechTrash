package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"techTrash/connection"
	"time"
)

func PostLog(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respError := map[string]string{"message": "could not read body"}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	var loglixeira []LogLixeira
	err = json.Unmarshal(body, &loglixeira)
	if err != nil {
		respError := map[string]string{"message": "could not unmarshal body"}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}
	idlixeira := loglixeira[0].Idlixeira
	nivel := loglixeira[0].Nivel
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")

	db, err := connection.MysqlConnect()
	if err != nil {
		respError := map[string]string{"message": "mysql failed to connect"}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO loglixeira (idlixeira, nivel, data) VALUES (%v, %v, "%v")`, idlixeira, nivel, date)
	log.Print(querySQL)
	_, err = db.Query(querySQL)
	if err != nil {
		respError := map[string]string{"message": "sql query failed to execute", "query": querySQL}
		jsonResp, _ := json.Marshal(respError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	if err == nil {
		resp := map[string]string{"message": "success"}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		return
	}

}
