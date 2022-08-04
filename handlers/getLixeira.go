package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"techTrash/connection"
)

type Lixeira struct {
	id          int     `json:"id"`
	localizacao string  `json:"localizacao"`
	nivel       float64 `json:"nivel"`
}

func GetLixeira(w http.ResponseWriter, r *http.Request) {
	db, err := connection.MysqlConnect()
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM lixeira")
	if err != nil {
		panic(err.Error())
	}
	var lixeira Lixeira
	for results.Next() {
		err = results.Scan(&lixeira.id, &lixeira.localizacao, &lixeira.nivel)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("Chesquedele")
	json.NewEncoder(w).Encode(lixeira)
}
