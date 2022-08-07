package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"techTrash/connection"
)

func PostLixeira(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}

	var lixeira []Lixeira
	err = json.Unmarshal(body, &lixeira)
	if err != nil {
		log.Print(err)
	}
	localizacao := lixeira[0].Localizacao

	db, err := connection.MysqlConnect()
	if err != nil {
		log.Print(ErrMysqlConnection)
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO lixeira (localizacao) VALUES ("%s")`, localizacao)
	log.Print(querySQL)
	_, err = db.Query(querySQL)
	if err != nil {
		log.Print(err)
	}

}
