package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func health() error {
	router := httprouter.New()
	router.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		_, _ = w.Write([]byte("hello xinyu"))
	})

	return http.ListenAndServe(":8080", router)
}
