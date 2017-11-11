package main

import (
	"fmt"
	"net/http"
	"html/template"

	"web/models"
	)

func indexHendler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)

}

func writeHendler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)

}

func savePostHendler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

}

func main()  {

	fmt.Println("Listening on port: 3000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/write", writeHendler)
	http.HandleFunc("/SavePost", savePostHendler)

	http.ListenAndServe(":3000", nil)

}
