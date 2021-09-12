package core

import (
	. "lol-api/api/v1/handlers"

	log "github.com/sirupsen/logrus"

	"net/http"
	"net/http/pprof"
)

const AV1 = "/api/v1"

func (s *Server) CreateRoutes() {

	s.User.Handle("/contact/mail", http.HandlerFunc(SendContactMail)).Methods("POST")

	s.User.Handle("/duo", http.HandlerFunc(GetDuoList)).Methods("GET")
	s.User.Handle("/duo", http.HandlerFunc(CreateDuo)).Methods("POST")
	s.User.Handle("/duo/{id:.*}", http.HandlerFunc(DeleteDuoByID)).Methods("DELETE")

	s.User.Handle("/test", http.HandlerFunc(TestGetAllColl)).Methods("GET")
	s.User.Handle("/test", http.HandlerFunc(TestCreateCollById)).Methods("POST")
	s.User.Handle("/test/{id:.*}", http.HandlerFunc(TestGetCollById)).Methods("GET")
	s.User.Handle("/test/{id:.*}", http.HandlerFunc(TestUpdateCollById)).Methods("PATCH")
	s.User.Handle("/test/{id:.*}", http.HandlerFunc(TestDeleteCollById)).Methods("DELETE")

	s.Admin.HandleFunc("/debug/pprof/", pprof.Index)
	s.Admin.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.Admin.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.Admin.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.Admin.HandleFunc("/debug/pprof/trace", pprof.Trace)

	s.Admin.Handle("/debug/pprof/block", pprof.Handler("block"))
	s.Admin.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	s.Admin.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	s.Admin.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	log.Info("------------------------------------------------")
}
