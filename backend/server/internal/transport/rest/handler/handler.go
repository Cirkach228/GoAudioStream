package handler

import ("fmt"
		"net/http"
)
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func runHandlers(){
	http.HandleFunc("/",mainPageHandler)
	http.ListenAndServe(":8080",nil)
}
