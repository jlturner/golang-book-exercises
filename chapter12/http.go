package main
import (
	"net/http"
	"io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(
		res,
		"<html><body><h1>Lady, Queen of French Africa</h1></body></html>",
	)
}

func main() {
	http.HandleFunc("/hello", hello)

	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	
	http.ListenAndServe(":9000", nil)
}
	
