package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	lg "github.com/GolangToolKits/go-level-logger"
	mux "github.com/GolangToolKits/grrt"
	hd "github.com/Ulbora/go-micro-blog-ui/handlers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func main() {
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var signinMap = make(map[string]s.Signin)
	var sh hd.MCHandler
	sh.Log = log
	// sh.Signins = signinMap

	//var linkedInClientID string
	//var linkedInClientSecret string

	//add signins here
	if os.Getenv("LINKED_IN_CLIENT_ID") != "" {
		linkedInClientID := os.Getenv("LINKED_IN_CLIENT_ID")
		linkedInClientSecret := os.Getenv("LINKED_IN_CLIENT_SECRET")
		var lsgn s.LinkedInSignin
		lsgn.Log = log
		lsgn.ClientID = linkedInClientID
		lsgn.ClientSecret = linkedInClientSecret
		signinMap["linkedin"] = lsgn.New()
		sh.Log.Info("clientId: ", linkedInClientID)
	}

	sh.Signins = signinMap

	// sh.Templates = template.Must(template.ParseFiles("./static/index.html",
	// 	"./static/product.html", "./static/addProduct.html"))

	router := mux.NewRouter()

	h := sh.New()

	//router.HandleFunc("/", h.Index).Methods("GET")
	//router.HandleFunc("/product/{id}/{sku}", h.ViewProduct).Methods("GET")
	//router.HandleFunc("/addProduct", h.AddProduct).Methods("POST")

	router.HandleFunc("/auth/linkedin/callback", h.LinkedInCallback).Methods("GET")

	port := "8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Web UI is running on port 8080!")

	http.ListenAndServe(":"+port, (router))
}

// go mod init github.com/Ulbora/go-micro-blog-ui
