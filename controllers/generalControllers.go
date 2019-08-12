package controllers

import "net/http"

var WelcomeController = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenido a el sistema GMDS."))
}
