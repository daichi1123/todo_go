package router

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, r)
}
