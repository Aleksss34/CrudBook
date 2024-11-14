package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Handlers() *mux.Router {
	rout := mux.NewRouter()
	rout.HandleFunc("/", mainPackage)
	return rout

}
func mainPackage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is main package"))
}
