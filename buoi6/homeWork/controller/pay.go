package controller

import (
	"fmt"
	"net/http"
)

func PayPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("token", "true")
	fmt.Fprintf(w, "pay page")
}
