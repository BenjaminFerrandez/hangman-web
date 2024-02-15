package main

import (
	"html/template"
	"net/http"
)

type Web struct {
	Word string
	Try  int
}

func Index(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/index.html"))
	template.Execute(w, nil)
	return
}

func Page2(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/page2.html"))
	template.Execute(w, nil)
	return
}

func Page3(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/page3.html"))
	template.Execute(w, nil)
	return
}

func PageJouer(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/pageJouer.html"))
	template.Execute(w, nil)
	return
}

func PageJugar(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/pageJugar.html"))
	template.Execute(w, nil)
	return
}

func PagePlay(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/pagePlay.html"))
	template.Execute(w, nil)
	return
}

func PageUrar(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/pageUrar.html"))
	template.Execute(w, nil)
	return
}

func GameGo(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("game.go"))
	template.Execute(w, nil)
	return
}

func Last(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/last.html"))
	r.ParseForm()
	letter := r.Form.Get("letter")
	InWord(letter)
	azerty := Web{
		Word: letter,
		Try:  10,
	}
	template.Execute(w, azerty)
	return
}

// maybe    if len(letter) > 0
//
