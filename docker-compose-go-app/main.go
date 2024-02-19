package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "https://www.google.com", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("You have a session cookie"))
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
