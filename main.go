package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//go through the files in the html folder
		//based on the url being checked for
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		var fileList []string
		var dirList []string
		for _, f := range files {
			if(f.IsDir()) {
				dirList = append(dirList,f.Name())
			} else {
				fileList = append(fileList,f.Name())
			}
		}
		fmt.Println("Files:",fileList)
		fmt.Println("Directorys:",dirList)
		//now parse the url being passed in
		spl := strings.Split(r.URL.String(),"/")
		
		if(spl[0] == "") {
			//serve the home page
		} else if(contains(dirList,spl[0]) && len(spl) == 3) {
			//serve the file in spl[1]
		} else if(contains(fileList,spl[0])) {
			//serve the file in spl[0]
		} else {
			//serve the 404 page
		}
		fmt.Println(spl)
		
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
	   if a == str {
		  return true
	   }
	}
	return false
 }
  

type server struct {}

func (s *server) serve() {
	log.Fatalln(http.ListenAndServe(":420", nil))
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