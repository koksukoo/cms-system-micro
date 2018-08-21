package main

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	hashids "github.com/speps/go-hashids"
	"golang.org/x/crypto/bcrypt"
)

// Credentials represent json body in post request
type Credentials struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
	ID       int    `json:"-" db:"id"`
}

var hashid = hashids.HashID{}

func init() {
	hd := hashids.NewData()
	hd.Salt = conf.Hashids.Salt
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)
	hashid = *h
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
		return
	}

	stmt, err := db.Prepare("SELECT * FROM users WHERE username LIKE (?)")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := stmt.QueryRow(credentials.Username)

	var dbCredentials Credentials
	err = res.Scan(&dbCredentials.ID, &dbCredentials.Username, &dbCredentials.Password)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbCredentials.Password), []byte(credentials.Password)); err != nil {
		respondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	sessionToken := uuid.NewV4()
	hash, _ := hashid.Encode([]int{dbCredentials.ID})
	// other services can access user hash from redis
	_, err = cache.Do("SETEX", sessionToken.String(), "86400", hash)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "cms_session",
		Path:    "/",
		Value:   sessionToken.String(),
		Expires: time.Now().AddDate(0, 0, 1),
	})

	respondJSON(w, http.StatusOK, map[string]string{"token": sessionToken.String()})
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
