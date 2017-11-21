package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sys-cat/linelogin"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("access index")
	fmt.Fprintf(w, "<h1>This is sys-cat test site</h1>")
}

func line_login(w http.ResponseWriter, r *http.Request) {
	log.Println("access line")
	urlParam := linelogin.New()
	err := urlParam.Parameters(os.Getenv("CHANNEL_ID"), os.Getenv("CHANNEL_SECRET"), os.Getenv("REDIRECT_URL"))
	if err != nil {
		data := fmt.Sprintf("<p>missing render: %s</p>", err.Error)
		fmt.Fprintf(w, data)
	}
	//fmt.Fprintf(w, urlParam.OutputURL())
	http.Redirect(w, r, urlParam.OutputURL(), 301)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
}

func main() {
	log.Println("------start server-----")
	http.HandleFunc("/", index)
	http.HandleFunc("/line_login", line_login)
	http.HandleFunc("/redirect", redirect)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
