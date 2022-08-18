package controllers

import (
	"fmt"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

func DeleteLixeira(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	query := r.URL.Query()
	var id []string
	id, ok := query["idlixeira"]
	if !ok || len(id) < 1 {
		utils.SetResponseError(w, r, "invalid id")
		return
	}
	idpassado := id[0]

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf("DELETE FROM lixeira WHERE idlixeira = %v", idpassado)
	_, err = db.Query("DELETE FROM lixeira WHERE idlixeira = ?", idpassado)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
