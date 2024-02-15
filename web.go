package main

import (
	"fmt"
	"net/http"
)

func web() {
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("font"))))

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	http.Handle("/son/", http.StripPrefix("/son/", http.FileServer(http.Dir("son"))))

	http.HandleFunc("/", Index)
	http.HandleFunc("/last", Last)
	http.HandleFunc("/index.html", Index)
	http.HandleFunc("/page2.html", Page2)
	http.HandleFunc("/page3.html", Page3)
	http.HandleFunc("/pageJouer.html", PageJouer)
	http.HandleFunc("/pageJugar.html", PageJugar)
	http.HandleFunc("/pagePlay.html", PagePlay)
	http.HandleFunc("/pageUrar.html", PageUrar)

	fmt.Println("(http://localhost:8080) - Server started on port ", PORT)
	http.ListenAndServe(":8080", nil)
}
