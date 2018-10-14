package main

import (
	"html/template"
	"net/http"
	"fmt"
	"os"
)

var (
	tpl            *template.Template
	DB             = dbConn()
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}


func port() string {
	port := os.Getenv("PORT")
	port = ":" + port
	return port
}


func Routes() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/projects", Projects)
	http.HandleFunc("/tutorials", Tutorials)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println(err)
	}
}
