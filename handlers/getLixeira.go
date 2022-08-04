package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"techTrash/connection"
	"errors"
)
var(
	ErrMysqlConnection = errors.New("could not connet to mysql")
)

type Lixeira struct {
	ID          int     `json:"id"`
	Localizacao string  `json:"localizacao"`
	Nivel       float64 `json:"nivel"`
}

func GetLixeira(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := connection.MysqlConnect()
	if err != nil {
		log.Print(ErrMysqlConnection)
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM lixeira")
	if err != nil {
		panic(err.Error())
	}
	var lixeira []Lixeira
	for results.Next() {
		var lixeirabanco Lixeira
		err = results.Scan(&lixeirabanco.ID, &lixeirabanco.Localizacao, &lixeirabanco.Nivel)
		if err != nil {
			panic(err.Error())
		}
		lixeira = append(lixeira, lixeirabanco)
	}
	json.NewEncoder(w).Encode(lixeira)
}
