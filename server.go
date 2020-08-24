package main

import (
	"log"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

type Bookmark struct {
	Name string
	Href string
}

type Collection struct {
	Name string
	Bookmarks []Bookmark
}

type Bookmarks struct {
	Collections []Collection
}

type PageData struct {
	Title string
	Bookmarks Bookmarks
}

func loadDataFromFile(path string) Bookmarks {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var bookmarks Bookmarks
	err = json.Unmarshal(data, &bookmarks)
	if err != nil {
		log.Println(err)
	}

	return bookmarks
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))

	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		log.Println("OK")
	})

	r.PathPrefix("/static/").
	  Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/bookmarks", func (w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("template/index.html"))

		bookmarks := loadDataFromFile("./data/test_data.json")
		pageData := PageData{
			Title: "Bookmarks",
			Bookmarks: bookmarks,
		}

		templ.Execute(w, pageData)
	})

	http.ListenAndServe(":4040", r)
}