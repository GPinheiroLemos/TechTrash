package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	query := r.URL.Query()
	id, ok := query["id"]
	if !ok || len(id) == 0 {
		log.Print(ErrMissingID)
	}
	idpassado := id[0]

	db, err := connection.MysqlConnect()
	if err != nil {
		log.Print(ErrMysqlConnection)
	}
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT * FROM lixeira WHERE idlixeira = %v", idpassado)
	results, err := db.Query(querySQL)
	if err != nil {
		panic(err.Error())
	}

	var lixeira []Lixeira
	for results.Next() {
		var lixeirabanco Lixeira
		err = results.Scan(&lixeirabanco.Id, &lixeirabanco.Localizacao)
		if err != nil {
			panic(err.Error())
		}
		lixeira = append(lixeira, lixeirabanco)
	}

	json.NewEncoder(w).Encode(lixeira)
}
