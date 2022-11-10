package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

var (
	ErrMysqlConnection = errors.New("could not connet to mysql")
	ErrMissingID       = errors.New("missing ID")
)

type Lixeira struct {
	Id          int    `json:"id"`
	Localizacao string `json:"localizacao"`
	Altura      string `json:"altura"`
}

func GetLixeira(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	var id []string
	id, ok := query["idlixeira"]
	if !ok || len(id) < 1 {
		id = append(id, "0")
	}
	idpassado := id[0]

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	var results *sql.Rows
	if idpassado == "0" {
		querySQL := "SELECT * FROM lixeira"
		results, err = db.Query("SELECT * FROM lixeira")
		if err != nil {
			message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
			utils.SetResponseError(w, r, message)
			return
		}
	} else {
		querySQL := fmt.Sprintf("SELECT * FROM lixeira WHERE idlixeira = %v", idpassado)
		results, err = db.Query("SELECT * FROM lixeira WHERE idlixeira = ?", idpassado)
		if err != nil {
			message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
			utils.SetResponseError(w, r, message)
			return
		}
	}

	var lixeira []Lixeira
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

	utils.SetResponseSuccess(w, r)
	json.NewEncoder(w).Encode(lixeira)

}
