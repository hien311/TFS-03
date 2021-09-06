package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello, this homepage")
}
