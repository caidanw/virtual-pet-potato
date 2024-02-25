package main

import (
	"io"
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", play)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var users = map[string]string{
	"test": "notsecret",
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}

	return password == pass
}

func play(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		http.RedirectHandler("/register", 401)
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		http.RedirectHandler("/login", 401)
		return
	}

	http.FileServer(http.Dir("./static/")).ServeHTTP(w, r)
}

func auth() {
}

func login(w http.ResponseWriter, r *http.Request) {
}

func register(w http.ResponseWriter, r *http.Request) {

}