package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello")
		d, err := io.ReadAll(r.Body)
		if err!=nil{
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("OOPS"))
			// can use the above code or use the below alternate
			http.Error(w, "OOPS", http.StatusBadRequest)
			return
		}

		// log.Printf("Data: %s", d)
		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye")
	})

	http.ListenAndServe(":9090",nil)
}