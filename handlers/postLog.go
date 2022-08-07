package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"techTrash/connection"
	"time"
)

func PostLog(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}

	var loglixeira []Log
	err = json.Unmarshal(body, &loglixeira)
	if err != nil {
		log.Print(err)
	}
	idlixeira := loglixeira[0].Idlixeira
	nivel := loglixeira[0].Nivel
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")

	db, err := connection.MysqlConnect()
	if err != nil {
		log.Print(ErrMysqlConnection)
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO loglixeira (idlixeira, nivel, data) VALUES (%v, %v, "%v")`, idlixeira, nivel, date)
	log.Print(querySQL)
	_, err = db.Query(querySQL)
	if err != nil {
		log.Print(err)
	}

}
