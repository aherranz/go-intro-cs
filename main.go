package main

import "fmt"
import "net/http"
import "encoding/json"

type Scope struct {
        Project string
        Area    string
}

type Note struct {
        Title string
        Tags  []string
        Text  string
        Scope Scope
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Note: %+v", note)
}

func main() {
	const serverAddr string = "127.0.0.1:8081"

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP Caracola\n"))
	})

	http.HandleFunc("POST /", createNote)
	
	http.ListenAndServe(serverAddr, nil)
	fmt.Println("Hola Caracola")
}
