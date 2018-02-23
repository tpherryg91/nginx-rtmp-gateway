package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const redirectRTMP = "rtmp://0.0.0.0:1935/live360p/%d"

func AuthHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	streamKeyIdentifier := "name"

	key := req.FormValue(streamKeyIdentifier)
	keyPort := make(map[string]int)
	keyPort["abcd"] = 8888
	keyPort["efgh"] = 8889

	if port, ok := keyPort[key]; ok {
		/*Set Port is Used*/
		http.Redirect(w, req, fmt.Sprintf(redirectRTMP, port), 301)
	} else {
		log.Println("["+strconv.Itoa(http.StatusForbidden)+"]", key)
		w.WriteHeader(http.StatusForbidden)
	}
}
