package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "google.com", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Welcome! Your session ID is: %s", sessionCookie.Value)
	})

	//Write a handler to redirect it to Google for now

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: "abc123", // You should generate a unique session ID here
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)

	//    http.HandleFunc("/", HelloServer)
	//    http.ListenAndServe("0.0.0.0:8080", nil)
	//}
	//
	//func HelloServer(w http.ResponseWriter, r *http.Request) {
	//    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
