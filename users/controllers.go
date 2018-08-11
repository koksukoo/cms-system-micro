package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Credentials represent json body in post request
type Credentials struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

func registerController(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)

	if _, err = db.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		credentials.Username,
		string(hashed),
	); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, credentials)
	return
}

func loginController(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	res := db.QueryRow("SELECT password FROM users WHERE username LIKE ?", credentials.Username)

	var dbCredentials Credentials
	err = res.Scan(&dbCredentials.Password)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbCredentials.Password), []byte(credentials.Password)); err != nil {
		respondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// there we would send jwt token back
	respondJSON(w, http.StatusOK, map[string]string{"token": "123123"})
}

func respondError(w http.ResponseWriter, code int, msg string) {
	respondJSON(w, code, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(code)
	w.Write(response)
}
