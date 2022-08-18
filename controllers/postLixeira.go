package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

func PostLixeira(w http.ResponseWriter, r *http.Request) {
	log.Print("Chegou uma requisição")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

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
	localizacao := lixeira[0].Localizacao

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO lixeira (localizacao) VALUES ("%s")`, localizacao)
	_, err = db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
