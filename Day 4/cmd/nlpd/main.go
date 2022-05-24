package main

import (
	"encoding/json"
	"fmt"
	"git.shipt.com/nlp"
	"io"
	"log"
	"net/http"
)

func main() {
	//routing
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/home", homeHandler)

	//run server
	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func tokeninzeHandler(w http.ResponseWriter, r *http.Request) {
	// get, parse, and validate input
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error: can't read - %s", err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	//work on data
	text := string(data)
	tokens := nlp.Tokenize(text)

	//marshal output
	resp := map[string]any{ //map is like a json object
		//string : any (empty interface)
		"tokens": tokens,
	}

	///* Option 1: Use json.Marshal
	//	+: Can check for marshal error
	//	-: Create data in memory
	out, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error: can't marshal - %s", err)
		http.Error(w, "can't Marshal", http.StatusInternalServerError)
		return
	}
	w.Write(out)

	//option 2: json encoder
	//+: No in memory, can marshal several values
	//-: Can't send error to client
	//if err := json.NewEncoder(w).Encode(resp); err != nil {
	//	//can't send error, only log
	//	log.Printf("error: can't marshal - %s", err)
	//}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<b>HOME</b>")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}
