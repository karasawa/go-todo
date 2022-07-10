package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	genelateHTML(w, "Hello", "layout", "public_navbar", "top")
}