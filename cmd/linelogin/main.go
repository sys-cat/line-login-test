package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sys-cat/linelogin"
	"github.com/sys-cat/linelogin/profile"
	"github.com/sys-cat/linelogin/token"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Println("access index")
	fmt.Fprintf(w, "<h1>This is sys-cat test site</h1>")
	url := linelogin.New()
	err := url.Parameters(os.Getenv("CHANNEL_ID"), os.Getenv("CHANNEL_SECRET"), os.Getenv("REDIRECT_URL"))
	if err != nil {
		fmt.Fprintf(w, "<p style=\"color:red;\">link url build error !</p>")
	}
	log.Printf("%+v\n", url.OutputURL())
	link := fmt.Sprintf("<a href=\"%s\">Line Login</a>", url.OutputURL())
	io.WriteString(w, link)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Println("access redirect")
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	if state == "" {
		fmt.Fprintf(w, "Invalid access\n")
	}
	newToken := token.New()
	err := newToken.Parameters(code, os.Getenv("REDIRECT_URL"), os.Getenv("CHANNEL_ID"), os.Getenv("CHANNEL_SECRET"))
	if err != nil {
		fmt.Fprintf(w, "Invalid parameters\n")
	}
	res, err := token.GetToken(newToken)
	if err != nil {
		fmt.Fprintf(w, "Get Token miss %s\n", err)
	}
	log.Println(fmt.Sprintf("id_token: %s", res.IDToken))
	profile, err := profile.GetProfileData(res.AccessToken)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("<h4>Get Profile miss %s</h4>\n", err))
	}
	if err == nil {
		io.WriteString(w, fmt.Sprintf("<img src=\"%s/small\" alt=\"profile image\">\n", profile.PictureURL))
		list := fmt.Sprintf("<ul>\n\t<li>ID: %s</li>\n\t<li>Name: %s</li>\n\t<li>Message: %s</li>\n</ul>\n", profile.UserID, profile.DisplayName, profile.StatusMessage)
		io.WriteString(w, list)
	}
}

func main() {
	log.Println("------start server-----")
	http.HandleFunc("/index", index)
	http.HandleFunc("/redirect", redirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
