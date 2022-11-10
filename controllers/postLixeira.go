package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"techTrash/connection"
	"techTrash/utils"
)

func PostLixeira(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
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
	altura, _ := strconv.ParseFloat(lixeira[0].Altura, 64)
	log.Print(altura)

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO lixeira (localizacao, altura) VALUES ("%s", %v)`, localizacao, altura)
	_, err = db.Query(querySQL)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
