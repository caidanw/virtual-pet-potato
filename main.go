package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu          sync.Mutex
	credentials = map[string]string{
		"username1": "password1",
		"username2": "password2",
	}
)

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "./static/auth.html")
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Access form data
	action := r.Form.Get("action")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	// fmt.Printf("A: %s\nU: %s\nP: %s\n", action, username, password)

	mu.Lock()
	defer mu.Unlock()

	if action == "register" {
		// Check if the username already exists
		if _, exists := credentials[username]; exists {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		// Register the new user
		credentials[username] = password

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		// Check credentials
		storedPassword, ok := credentials[username]
		if !ok || storedPassword != password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/auth", auth)

	fmt.Println("Server running on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
