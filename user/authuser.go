package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SetResponseError(w, r, "could not read body")
		return
	}

	var newuser []User
	err = json.Unmarshal(body, &newuser)
	if err != nil {
		utils.SetResponseError(w, r, "could not unmarshal body")
		return
	}

	username := newuser[0].Username
	password := newuser[0].Password

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`SELECT password FROM usuario WHERE username = %s`, username)
	results, err := db.Query("SELECT password FROM usuario WHERE username = ?", username)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	var usuario []User
	for results.Next() {
		var usuarioTechTrash User
		err = results.Scan(&usuarioTechTrash.Password)
		if err != nil {
			utils.SetResponseError(w, r, "error while reading data response from database")
			return
		}
		usuario = append(usuario, usuarioTechTrash)
	}

	auth := utils.CheckPasswordHash(password, usuario[0].Password)
	if auth == false {
		utils.SetResponseError(w, r, "wrong password")
		return
	}

	utils.SetResponseSuccess(w, r)

}
