package main

import (
	"github.com/julienschmidt/httprouter"
	grace "gopkg.in/paytm/grace.v1"
)

func main() {
	router := httprouter.New()
	router.POST("/auth", AuthHandler)
	grace.Serve(":5000", router)
}
