package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routing
	// Rules: If route ends with / - prefix match, otherwise - exact match
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/home", homeHandler)
	//http.HandleFunc("/tokenize", tokenizeHandler)
	http.HandleFunc("/stem/{word}", stemHandler) //does not work?????

	//run server
	addr := ":8080"
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL

	fmt.Fprintln(w, v)
}

//func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
//	// get, parse, and validate input
//	defer r.Body.Close()
//
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Printf("error: can't read - %s", err)
//		http.Error(w, "can't read", http.StatusBadRequest)
//		return
//	}
//
//	//work on data
//	text := string(data)
//	tokens := nlp.Tokenize(text)
//
//	//marshal output
//	resp := map[string]any{ //map is like a json object
//		//string : any (empty interface)
//		"tokens": tokens,
//	}
//
//	///* Option 1: Use json.Marshal
//	//	+: Can check for marshal error
//	//	-: Create data in memory
//	out, err := json.Marshal(resp)
//	if err != nil {
//		log.Printf("error: can't marshal - %s", err)
//		http.Error(w, "can't Marshal", http.StatusInternalServerError)
//		return
//	}
//	w.Write(out)
//
//	//option 2: json encoder
//	//+: No in memory, can marshal several values
//	//-: Can't send error to client
//	//if err := json.NewEncoder(w).Encode(resp); err != nil {
//	//	//can't send error, only log
//	//	log.Printf("error: can't marshal - %s", err)
//	//}
//}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<b>HOME</b>")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

/*
type Server struct {
	db *DB
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	db, err := connectDB(dsn)
	...
	s =  Server{db: db} // dependency injection

	http.HandleFunc("/health", s.healthHandler)
}
*/
