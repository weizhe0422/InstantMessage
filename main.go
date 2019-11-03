package main

import (
	"html/template"
	"log"
	"net/http"

	"./Controller"
)





func registerView() {
	tpls, err := template.ParseGlob("View/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpls.Templates() {
		tplName := v.Name()
		log.Println(tplName)
		http.HandleFunc(tplName, func(w http.ResponseWriter, r *http.Request) {
			tpls.ExecuteTemplate(w, tplName, nil)
		})

	}
}

func main() {

	http.HandleFunc("/user/login", Controller.LoginFunc())
	http.HandleFunc("/user/register", Controller.RegisterFunc())
	http.Handle("/", http.FileServer(http.Dir(".")))
	/*http.HandleFunc("/user/login.shtml",func(w http.ResponseWriter, r *http.Request){
		tpl, err := template.ParseFiles("View/user/login.html")
		if err != nil{
			log.Fatal(err.Error())
		}
		tpl.ExecuteTemplate(w, "/user/login.shtml", nil)
	}*/
	registerView()
	http.ListenAndServe(":8080", nil)
}




