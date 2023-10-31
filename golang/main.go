package main

import (
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, no-transform")

	var tmpl, err = template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func providerPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, no-transform")
	w.Header().Add("Location", "/callback")
	w.WriteHeader(http.StatusPermanentRedirect)
}

func callbackPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, no-transform")
	w.Header().Add("Set-Cookie", "session=nedya-amril-prakasa-golang")
	w.WriteHeader(http.StatusOK)

	var tmpl, err = template.ParseFiles("./callback.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/provider", providerPage)
	http.HandleFunc("/callback", callbackPage)

	http.ListenAndServe(":3000", nil)
}
