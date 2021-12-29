package handler

import (
	"belajar/entity"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := entity.Data{Description: "Welcome Home"}

	tmpl, err := template.ParseFiles(path.Join("view", "Home.html"), path.Join("view", "Layout.html"))

	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	numberId, err := strconv.Atoi(id)

	if err != nil || numberId <= 0 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("view", "Product.html"), path.Join("view", "Layout.html"))

	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	data := entity.Product{
		ID: int64(numberId),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("view", "Store.html"), path.Join("view", "Layout.html"))

	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	data := []entity.Store{
		{ID: 1, Name: "Store 1", Stock: 13},
		{ID: 2, Name: "Store 2", Stock: 2},
		{ID: 3, Name: "Store 3", Stock: 9},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error happens", http.StatusInternalServerError)
		log.Println(err.Error())
	}

}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("view", "Form.html"), path.Join("view", "Layout.html"))

		if err != nil {
			http.Error(w, "Error happens", http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error happens", http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}
	}
	http.Error(w, "Error happens", http.StatusInternalServerError)

}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error happens", http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		templ, err := template.ParseFiles(path.Join("view", "Process.html"), path.Join("view", "Layout.html"))

		if err != nil {
			http.Error(w, "Error happens", http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		data := entity.Formdata{
			Name:    r.Form.Get("name"),
			Message: r.Form.Get("message"),
		}

		log.Println(data.Name)
		log.Println(data.Message)
		templ.Execute(w, data)

		if err != nil {
			http.Error(w, "Error happens", http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		return
	}
	http.Error(w, "Error happens", http.StatusInternalServerError)
}
