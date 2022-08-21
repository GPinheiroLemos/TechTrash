package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"techTrash/connection"
	"techTrash/utils"
)

func EditLixeira(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SetResponseError(w, r, "could not read body")
		return
	}

	var lixeira []Lixeira
	err = json.Unmarshal(body, &lixeira)
	if err != nil {
		utils.SetResponseError(w, r, "could not unmarshal body")
		return
	}
	id := lixeira[0].Id
	localizacao := lixeira[0].Localizacao
	altura, _ := strconv.ParseFloat(lixeira[0].Altura, 64)

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT * FROM lixeira WHERE idlixeira = %v", id)
	results, err := db.Query("SELECT * FROM lixeira WHERE idlixeira = ?", id)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	for results.Next() {
		var lixeirabanco Lixeira
		err = results.Scan(&lixeirabanco.Id, &lixeirabanco.Localizacao, &lixeirabanco.Altura)
		if err != nil {
			respError := map[string]string{"message": "error while reading data response from database"}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
		lixeira = append(lixeira, lixeirabanco)
	}

	if altura == 0 {
		altura, _ = strconv.ParseFloat(lixeira[0].Altura, 64)
	}
	if localizacao == "" {
		localizacao = lixeira[0].Localizacao
	}

	querySQL = fmt.Sprintf(`UPDATE lixeira SET localizacao = "%s", altura = %v WHERE idlixeira = %d`, localizacao, altura, id)
	_, err = db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
