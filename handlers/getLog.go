package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"techTrash/connection"
)

type Log struct {
	Idlog     int     `json:"idlog"`
	Idlixeira int     `json:"idlixeira"`
	Nivel     float64 `json:"nivel"`
	Data      string  `json:"data"`
}

func GetLog(w http.ResponseWriter, r *http.Request) {

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

	querySQL := fmt.Sprintf("SELECT * FROM loglixeira WHERE idlixeira = %v", idpassado)
	results, err := db.Query(querySQL)
	if err != nil {
		panic(err.Error())
	}

	var log []Log
	for results.Next() {
		var logbanco Log
		err = results.Scan(&logbanco.Idlog, &logbanco.Idlixeira, &logbanco.Nivel, &logbanco.Data)
		if err != nil {
			panic(err.Error())
		}
		log = append(log, logbanco)
	}

	json.NewEncoder(w).Encode(log)
}
