package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"techTrash/connection"
)

var (
	ErrMysqlConnection = errors.New("could not connet to mysql")
	ErrMissingID       = errors.New("missing ID")
)

type Lixeira struct {
	Id          int    `json:"id"`
	Localizacao string `json:"localizacao"`
}

func GetLixeira(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()
	var id []string
	id, ok := query["idlixeira"]
	if !ok || len(id) < 1 {
		id = append(id, "0")
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

	var results *sql.Rows
	if idpassado == "0" {
		querySQL := fmt.Sprintf("SELECT * FROM lixeira")
		results, err = db.Query("SELECT * FROM lixeira")
		if err != nil {
			respError := map[string]string{"message": "sql query failed to execute", "query": querySQL}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
	} else {
		querySQL := fmt.Sprintf("SELECT * FROM lixeira WHERE idlixeira = %v", idpassado)
		results, err = db.Query("SELECT * FROM lixeira WHERE idlixeira = ?", idpassado)
		if err != nil {
			respError := map[string]string{"message": "sql query failed to execute", "query": querySQL}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
	}

	var lixeira []Lixeira
	for results.Next() {
		var lixeirabanco Lixeira
		err = results.Scan(&lixeirabanco.Id, &lixeirabanco.Localizacao)
		if err != nil {
			respError := map[string]string{"message": "error while reading data response from database"}
			jsonResp, _ := json.Marshal(respError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResp)
			return
		}
		lixeira = append(lixeira, lixeirabanco)
	}

	if err == nil {
		resp := map[string]string{"message": "success"}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		json.NewEncoder(w).Encode(lixeira)
		return
	}
	return
}
