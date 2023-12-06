package main

import (
	"net/http"
	"strings"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

type route struct {
	handler   handler.APIFunc
	methods   string // comma separated
	modifiers string // comma separated
}

type router struct {
	r      *mux.Router
	stores db.StoreHolder
	routes map[string]route
}

func (r router) serve(addr string) {
	for k, v := range r.routes {
		r.r.HandleFunc(strings.TrimRight(k, " "), handler.MW(v.handler, r.stores, strings.Split(v.modifiers, ",")...)).Methods(strings.Split(v.methods, ",")...)
	}
	http.ListenAndServe(addr, r.r)
}
