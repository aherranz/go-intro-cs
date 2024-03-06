package main

import "fmt"

func main() {
	const serverAddr string = "127.0.0.1:4000"
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP Caracola"))
		http.ListenAndServe(serverAddr, nil)
	})
	fmt.Println("Hola Caracola")
}
