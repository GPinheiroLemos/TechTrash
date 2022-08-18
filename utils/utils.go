package utils

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SetResponseSuccess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	return
}

func SetResponseError(w http.ResponseWriter, r *http.Request, message string) {
	respSuccess := map[string]string{"message": message}
	jsonResp, _ := json.Marshal(respSuccess)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResp)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
