package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func AuthHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	key := req.FormValue("name")

	keyInt, err := strconv.Atoi(key)
	if err != nil {
		log.Println("["+strconv.Itoa(http.StatusForbidden)+"]", key, ":", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if keyInt >= 8888 && keyInt <= 8889 {
		w.WriteHeader(http.StatusOK)
	}
}
