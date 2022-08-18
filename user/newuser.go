package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"techTrash/connection"
	"techTrash/utils"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(w http.ResponseWriter, r *http.Request) {

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

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		utils.SetResponseError(w, r, "could not encrypt password")
		return
	}

	db, err := connection.MysqlConnect()
	if err != nil {
		utils.SetResponseError(w, r, "mysql failed to connect")
		return
	}
	defer db.Close()

	querySQL := fmt.Sprintf(`INSERT INTO usuario (username,password) VALUES (%s, %s)`, username, "********")
	_, err = db.Query("INSERT INTO usuario (username,password) VALUES (?, ?)", username, hashPassword)
	if err != nil {
		message := fmt.Sprintf("mysql query failed to execute. query: %s", querySQL)
		utils.SetResponseError(w, r, message)
		return
	}

	utils.SetResponseSuccess(w, r)

}
