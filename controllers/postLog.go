package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
	"time"
)

func PostLog(w http.ResponseWriter, r *http.Request) {

	// w.Header().Add("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SetResponseError(w, r, "could not read body")
		return
	}

	var loglixeira []LogLixeira
	err = json.Unmarshal(body, &loglixeira)
	if err != nil {
		utils.SetResponseError(w, r, "could not unmarshal body")
		return
	}
	idlixeira := loglixeira[0].Idlixeira
	nivel := loglixeira[0].Nivel
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO loglixeira (idlixeira, nivel, data) VALUES (%v, %v, "%v")`, idlixeira, nivel, date)
	_, err = db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
