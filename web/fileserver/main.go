package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/russross/blackfriday"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request){
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")));
	rw.Write(markdown)
}

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)

	http.Handle("/count", new(countHandler))

	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", nil)
}