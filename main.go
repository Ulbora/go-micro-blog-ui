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
	// var linkedInSigninRedirectURL string
	// var googleSigninRedirectURL string

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
		linkedInRedirectURL := os.Getenv("LINKED_IN_SIGNIN_REDIRECT_URL")
		var lsgn s.LinkedInSignin
		lsgn.Log = log
		lsgn.ClientID = linkedInClientID
		lsgn.ClientSecret = linkedInClientSecret
		if linkedInRedirectURL != "" {
			lsgn.RedirectURI = linkedInRedirectURL
		}
		signinMap["linkedIn"] = lsgn.New()
		sh.Log.Info("clientId: ", linkedInClientID)
	}
	if os.Getenv("GOOGLE_CLIENT_ID") != "" {
		googleOAuth2ClientID := os.Getenv("GOOGLE_CLIENT_ID")
		googleOAuth2ClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
		googleRedirectURL := os.Getenv("GOOGLE_SIGNIN_REDIRECT_URL")
		var glsgn s.GoogleSignin
		glsgn.Log = log
		glsgn.ClientID = googleOAuth2ClientID
		glsgn.ClientSecret = googleOAuth2ClientSecret
		if googleRedirectURL != "" {
			glsgn.RedirectURI = googleRedirectURL
		}
		signinMap["googleOAuth2"] = glsgn.New()
		sh.Log.Info("clientId: ", googleOAuth2ClientID)
	}

	sh.Signins = signinMap

	sh.Templates = template.Must(template.ParseFiles("./static/login.html",
		"./static/blogList.html", "./static/header.html", "./static/user-nav.html",
		"./static/addBlog.html", "./static/blog.html", "./static/editBlog.html",
		"./static/user-admin-nav.html"))

	sh.AdminTemplates = template.Must(template.ParseFiles("./static/admin/index.html",
		"./static/header.html", "./static/admin/adminBlogList.html",
		"./static/user-admin-nav.html", "./static/admin/adminBlog.html",
		"./static/admin/adminUser.html", "./static/admin/adminUnactivatedUser.html",
		"./static/admin/adminBannedUser.html", "./static/admin/adminConfig.html"))

	router := mux.NewRouter()

	h := sh.New()

	port := "8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.HandleFunc("/auth/linkedin/callback", h.LinkedInCallback).Methods("GET")
	router.HandleFunc("/signin-google/callback", h.GoogleSigninCallback).Methods("GET")
	router.HandleFunc("/login", h.LoginUserPage).Methods("GET")
	router.HandleFunc("/loginUser/{signingSystem}", h.LoginUser).Methods("GET")
	router.HandleFunc("/", h.GetBlogList).Methods("GET")
	router.HandleFunc("/newPost", h.AddBlogPage).Methods("GET")
	router.HandleFunc("/addPost", h.AddBlog).Methods("POST")
	router.HandleFunc("/addLike/{bid}", h.AddLike).Methods("GET")
	router.HandleFunc("/viewPost/{bid}", h.GetBlog).Methods("GET")
	router.HandleFunc("/addComment", h.AddComment).Methods("POST")
	router.HandleFunc("/editPost/{bid}", h.UpdateBlogPage).Methods("GET")
	router.HandleFunc("/updatePost", h.UpdateBlog).Methods("POST")

	router.HandleFunc("/adminPostList", h.GetAdminBlogList).Methods("GET")
	router.HandleFunc("/activateBlog/{bid}", h.ActivateBlog).Methods("GET")
	router.HandleFunc("/deactivateBlog/{bid}", h.DeactivateBlog).Methods("GET")
	router.HandleFunc("/deleteBlog/{bid}", h.DeleteBlog).Methods("GET")
	router.HandleFunc("/adminViewPost/{bid}", h.GetAdminCommentList).Methods("GET")
	router.HandleFunc("/activateComment/{cid}/{bid}", h.ActivateComment).Methods("GET")
	router.HandleFunc("/deactivateComment/{cid}/{bid}", h.DeactivateComment).Methods("GET")
	router.HandleFunc("/adminUser/{uid}", h.GetUserByIDPage).Methods("GET")
	router.HandleFunc("/adminInactiveUserList", h.GetUnactivatedUserList).Methods("GET")
	router.HandleFunc("/adminBannedUserList", h.GetBannedUserList).Methods("GET")
	router.HandleFunc("/activateUser/{uid}", h.EnableUser).Methods("GET")
	router.HandleFunc("/disableUser/{uid}", h.DisableUser).Methods("GET")
	router.HandleFunc("/deactivateUser/{uid}", h.DisableUserForCause).Methods("GET")
	router.HandleFunc("/reactivateUser/{uid}", h.ReinstateBannedUser).Methods("GET")
	router.HandleFunc("/adminConfig", h.GetConfig).Methods("GET")
	router.HandleFunc("/updateConfig", h.UpdateConfig).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Web UI is running on port 8080!")

	http.ListenAndServe(":"+port, (router))
}

// go mod init github.com/Ulbora/go-micro-blog-ui
