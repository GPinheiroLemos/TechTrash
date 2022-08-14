package utils

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SetResponseSuccess(w http.ResponseWriter, r *http.Request, message string) {
	respError := map[string]string{"message": message}
	jsonResp, _ := json.Marshal(respError)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
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
