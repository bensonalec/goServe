package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"html/template"

)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//go through the files in the html folder
		//based on the url being checked for
		files, err := ioutil.ReadDir("./html/")
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

		spl := strings.Split(r.URL.String(),"/")
		fmt.Println(spl)
		if(spl[1] == "") {
			//serve the home page
			fmt.Println("Served index")
			err = servePage(w,r,"index.html")

		} else if(contains(dirList,spl[1]) && len(spl) == 3) {
			//serve the file in spl[1]
			//check if spl[1] is a file inside the folder spl[0]
			innerFiles, err := ioutil.ReadDir("./html/"+spl[1]+"/")
			if err != nil {
				log.Fatal(err)
			}
	
			var innerFilesList []string
			for _,f := range innerFiles {
				if(!f.IsDir()){
					innerFilesList = append(innerFilesList,f.Name())
				}
				
			}
			if(contains(innerFilesList,spl[2]+".html")) {
				err = servePage(w,r,"/" + spl[1]+ "/" +  spl[2] + ".html")
				fmt.Println("Serving:","/" + spl[1]+ "/" +  spl[2] + ".html")
			} else {
				fmt.Println("Serving 404")
				http.Redirect(w,r,"/",http.StatusSeeOther)
			}
			
		} else if(contains(fileList,spl[1]+".html")) {
			//serve the file in spl[0]
			fmt.Println("Serving",spl[1])
			err = servePage(w,r,""+ spl[1] +".html")
		} else if(spl[1] == "favicon.ico") {

		} else {
			//serve the 404 page
			fmt.Println("Served 404")
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}
		if err != nil {
			fmt.Println("Error loading page")
			
		}
		
	}
}

func servePage(w http.ResponseWriter, r *http.Request, pageName string) error{
	t := template.Must(template.ParseFiles("html/"+pageName))
	err := t.Execute(w,nil)
	
	return err

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
	   if a == str {
		  return true
	   }
	}
	return false
 }
