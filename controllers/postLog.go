package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"techTrash/connection"
	"techTrash/utils"
	"time"
)

type LogLixeira struct {
	Idlog     int     `json:"idlog"`
	Idlixeira int     `json:"idlixeira"`
	Nivel     float64 `json:"nivel"`
	Data      string  `json:"data"`
	Distancia float64 `json:"distancia"`
}

func PostLog(w http.ResponseWriter, r *http.Request) {

	log.Print("Chegou uma resquisição do esp!")

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
	distancia := loglixeira[0].Distancia
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT * FROM lixeira WHERE idlixeira = %v", idlixeira)
	results, err := db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	var dadosLixeira []Lixeira
	for results.Next() {
		var lixeiraBanco Lixeira
		err = results.Scan(&lixeiraBanco.Id, &lixeiraBanco.Localizacao, &lixeiraBanco.Altura)
		if err != nil {
			respError := map[string]string{"message": "error while reading data response from database"}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
		dadosLixeira = append(dadosLixeira, lixeiraBanco)
	}
	altura, _ := strconv.ParseFloat(dadosLixeira[0].Altura, 64)

	querySQL = fmt.Sprintf("SELECT nivel FROM loglixeira WHERE idlixeira = %v", idlixeira)
	results, err = db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	var dadosLogLixeira []LogLixeira
	for results.Next() {
		var logLixeiraBanco LogLixeira
		err = results.Scan(&logLixeiraBanco.Nivel)
		if err != nil {
			respError := map[string]string{"message": "error while reading data response from database"}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
		dadosLogLixeira = append(dadosLogLixeira, logLixeiraBanco)
	}

	nivelAnterior := dadosLogLixeira[len(dadosLogLixeira)-1].Nivel

	nivel := ((distancia - 3 - altura) * (-1)) / altura * 100
	nivel = math.Round(nivel*100) / 100
	if nivel >= 100 {
		nivel = 100
	} else if nivel <= 0 {
		nivel = 0
	}

	diferenca := math.Abs(nivelAnterior - nivel)
	if diferenca > 3 {
		querySQL = fmt.Sprintf(`INSERT INTO loglixeira (idlixeira, nivel, data) VALUES (%v, %v, "%v")`, idlixeira, nivel, date)
		_, err = db.Query(querySQL)
		if err != nil {
			message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
			utils.SetResponseError(w, r, message)
			return
		}
	}

	utils.SetResponseSuccess(w, r)

}
