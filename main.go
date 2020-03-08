package main

import (
	"net/http"
	"log"

)

type server struct {}

func (s *server) serve() {
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func (s *server) routes() {
	http.HandleFunc("/", s.handleIndex())

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs)) 
	
}


func main() {
	serv := server{}
	serv.routes()
	serv.serve()
}