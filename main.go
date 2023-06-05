package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
	hd "github.com/Ulbora/go-micro-blog-ui/handlers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func main() {
	var secretSessionKey string

	var siteTitle string
	var siteDesc string
	var siteKeyWords string
	var restURL string
	var apiKey string
	var apiAdminKey string

	if os.Getenv("SECRET_SESSION_KEY") != "" {
		secretSessionKey = os.Getenv("SECRET_SESSION_KEY")
	} else {
		secretSessionKey = "dsdfsadfs61dsscfsdfdsdsfsdsdllsdfgggg"
	}

	if os.Getenv("SITE_TITLE") != "" {
		siteTitle = os.Getenv("SITE_TITLE")
	} else {
		siteTitle = "Go Micro Blog"
	}

	if os.Getenv("SITE_DESC") != "" {
		siteDesc = os.Getenv("SITE_DESC")
	} else {
		siteDesc = "A Micro Blog written in Golang by Ulbora Labs LLC"
	}

	if os.Getenv("SITE_KEY_WORDS") != "" {
		siteKeyWords = os.Getenv("SITE_KEY_WORDS")
	} else {
		siteKeyWords = "Golang, Micro, Blog, Ulbora, Labs"
	}

	if os.Getenv("REST_URL") != "" {
		restURL = os.Getenv("REST_URL")
	} else {
		restURL = "http://localhost:3000"
	}

	if os.Getenv("API_KEY") != "" {
		apiKey = os.Getenv("API_KEY")
	} else {
		apiKey = "557444414141"
	}

	if os.Getenv("API_ADMIN_KEY") != "" {
		apiAdminKey = os.Getenv("API_ADMIN_KEY")
	} else {
		apiAdminKey = "54211789991515"
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager(secretSessionKey, cf)
	if err != nil {
		log.Info(err)
		log.Info("Session err: ", err)
	}

	var signinMap = make(map[string]s.Signin)

	var del mcd.MCDelegate
	del.Log = log
	del.RestURL = restURL
	del.APIAdminKey = apiAdminKey
	del.APIKey = apiKey

	dd := del.New()

	var sh hd.MCHandler
	sh.Title = siteTitle
	sh.Desc = siteDesc
	sh.KeyWords = siteKeyWords
	sh.Log = log
	sh.SessionManager = sessionManager
	sh.Delegate = dd
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
		signinMap["linkedIn"] = lsgn.New()
		sh.Log.Info("clientId: ", linkedInClientID)
	}

	sh.Signins = signinMap

	sh.Templates = template.Must(template.ParseFiles("./static/login.html",
		"./static/blogList.html", "./static/header.html"))

	sh.AdminTemplates = template.Must(template.ParseFiles("./static/admin/index.html",
		"./static/header.html"))

	router := mux.NewRouter()

	h := sh.New()

	//router.HandleFunc("/", h.Index).Methods("GET")
	//router.HandleFunc("/product/{id}/{sku}", h.ViewProduct).Methods("GET")
	//router.HandleFunc("/addProduct", h.AddProduct).Methods("POST")

	// router.HandleFunc("/auth/linkedin/callback", h.LinkedInCallback).Methods("GET")

	port := "8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.HandleFunc("/auth/linkedin/callback", h.LinkedInCallback).Methods("GET")
	router.HandleFunc("/login", h.LoginUserPage).Methods("GET")
	router.HandleFunc("/loginUser/{signingSystem}", h.LoginUser).Methods("GET")
	router.HandleFunc("/", h.GetBlogList).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Web UI is running on port 8080!")

	http.ListenAndServe(":"+port, (router))
}

// go mod init github.com/Ulbora/go-micro-blog-ui
