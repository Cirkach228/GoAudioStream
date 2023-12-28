package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	//"os"
	"path/filepath"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request){
	gopath := os.Getenv("GOPATH")
	currentDir, _ := os.Getwd()  // Get the current working directory
	pathToIndexHtml := filepath.Join(currentDir, "frontend", "mainPage", "index.html")
	tmpl, err := template.ParseFiles("$GOPATH/GoAudioStream/frontend/mainPage/index.html") 
	if err != nil{ // catching parse error
		panic(err)
	}
	tmpl.Execute(w,"using template") // draw page
}

func RunHandlers(){
	mux := http.NewServeMux() // creating mux for requests
	mux.HandleFunc("/",mainPageHandler)// handle main page
	err := http.ListenAndServe(":8083",mux) // starting server
	if err != nil{ //catching start server error
		fmt.Println("%v",err)
	}
}
