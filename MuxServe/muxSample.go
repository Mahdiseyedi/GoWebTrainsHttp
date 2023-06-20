package main

import (
	"io"
	"net/http"
)

type MyHandler int

func (h MyHandler) firstHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "first Handler\t "+req.RequestURI+"\t "+req.Method+"\n")
}
func (h MyHandler) secondHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "second Handler\t "+req.RequestURI+"\t "+req.Method+"\n")
}
func (h MyHandler) defaultHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world\t "+req.RequestURI+"\t "+req.Method+"\n")
}

func main() {
	var h MyHandler
	mux := http.NewServeMux()
	mux.HandleFunc("/first", h.firstHandler)
	mux.HandleFunc("/second", h.secondHandler)
	mux.HandleFunc("/", h.defaultHandler)

	http.ListenAndServe(":9000", mux)

}
