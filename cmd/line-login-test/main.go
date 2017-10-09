package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/sys-cat/line-login-test"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln("<h1>This is sys-cat test site</h1>")
}

func line_login(w http.ResponseWriter, r *http.Request) {
	urlParam := line_login_test.New()
	if err := urlParam.Parameters(os.Getev("CHANNEL_ID"), os.Getev("CHANNE_SECRET"), os.Getev("REDIRECT_URL")); err != nil {
		fmt.Fprintln("<p>missing render: %s</p>", err.Error)
	}
	http.Redirect(w, r, urlParam.OutputURL(), 301)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/line", line_login)
	http.HandleFunc("/redirect", redirect)
	log.Fatal(http.ListenAndServe(":80", nil))
}
