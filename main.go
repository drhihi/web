package main

import (
	"fmt"
	"net/http"
	"html/template"

	"web/models"
)

var posts map[string]*models.Post

func indexHendler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", posts)

}

func writeHendler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)

}

func editHendler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	p, ok := posts[id]
	if !ok {
		http.NotFound(w, r)
	}
	t.ExecuteTemplate(w, "write", p)

}

func savePostHendler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = generateId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}

	http.Redirect(w, r,"/", 302)

}

func deleteHendler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	_, ok := posts[id]
	if !ok {
		http.NotFound(w, r)
	}
	delete(posts, id)

	http.Redirect(w, r,"/", 302)

}

func main()  {

	fmt.Println("Listening on port: 3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/write", writeHendler)
	http.HandleFunc("/edit", editHendler)
	http.HandleFunc("/delete", deleteHendler)
	http.HandleFunc("/SavePost", savePostHendler)

	http.ListenAndServe(":3000", nil)

}
