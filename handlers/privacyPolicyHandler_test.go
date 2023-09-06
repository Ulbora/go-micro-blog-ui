package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_SetPrivacyPolicyPage(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/", nil)
	vars := map[string]string{
		"uid": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("isAdmin", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("GET", "/", nil)

	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", true)
	ses2.Set("isAdmin", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var mcdel mcd.MockDelegate
	var rl mcd.PrivacyPolicy
	rl.ID = 2
	rl.Content = "test"
	mcdel.MockPrivacyPolicy = &rl

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		SiteName       string
		Title          string
		Desc           string
		KeyWords       string
		SiteURL        string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		resCode int
		w       *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				AdminTemplates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w,
				r: r,
			},
			resCode: 200,
			w:       w,
		},
		{
			name: "test 2",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				AdminTemplates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w2,
				r: r2,
			},
			resCode: 302,
			w:       w2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Delegate:       tt.fields.Delegate,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				SiteName:       tt.fields.SiteName,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
				SiteURL:        tt.fields.SiteURL,
			}
			h.SetPrivacyPolicyPage(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_SetPrivacyPolicy(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("POST", "/setPrivacyPolicy", strings.NewReader("content=some content"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("isAdmin", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("POST", "/setPrivacyPolicy", strings.NewReader("id=3&content=some content"))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", true)
	ses2.Set("isAdmin", true)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	r3, _ := http.NewRequest("POST", "/setPrivacyPolicy", strings.NewReader("content=some content"))
	r3.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w3 := httptest.NewRecorder()

	ses3 := sessionManager.NewSession(r3, "go-micro-blog-ui")
	ses3.Set("loggedIn", true)
	ses3.Set("isAdmin", false)
	ses3.Set("userEmail", "bob@bob.com")
	ses3.Save(w3)

	cook3 := w3.Result().Cookies()
	if len(cook3) > 0 {
		r3.AddCookie(cook3[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	mcdel.MockAddPrivacyPolicyResp = &mcd.ResponseID{
		ID:      3,
		Success: true,
	}
	mcdel.MockUpdatePrivacyPolicyResp = &mcd.Response{
		Success: false,
	}

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		SiteName       string
		Title          string
		Desc           string
		KeyWords       string
		SiteURL        string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		resCode int
		w       *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				AdminTemplates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w,
				r: r,
			},
			resCode: 302,
			w:       w,
		},
		{
			name: "test 2",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				AdminTemplates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w2,
				r: r2,
			},
			resCode: 302,
			w:       w2,
		},
		{
			name: "test 3",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				AdminTemplates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w3,
				r: r3,
			},
			resCode: 302,
			w:       w3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Delegate:       tt.fields.Delegate,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				SiteName:       tt.fields.SiteName,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
				SiteURL:        tt.fields.SiteURL,
			}
			h.SetPrivacyPolicy(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetPrivacyPolicy(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/", nil)
	// vars := map[string]string{
	// 	"uid": "5",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("isAdmin", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("GET", "/", nil)
	// vars := map[string]string{
	// 	"uid": "5",
	// }
	// r = mux.SetURLVars(r, vars)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("isAdmin", true)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var mcdel mcd.MockDelegate
	var rl mcd.PrivacyPolicy
	rl.ID = 2
	rl.Content = "test"
	mcdel.MockPrivacyPolicy = &rl

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		SiteName       string
		Title          string
		Desc           string
		KeyWords       string
		SiteURL        string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		resCode int
		w       *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				w: w,
				r: r,
			},
			resCode: 200,
			w:       w,
		},
		// {
		// 	name: "test 2",
		// 	fields: fields{
		// 		Log:            log,
		// 		Delegate:       &mcdel,
		// 		SessionManager: sessionManager,
		// 		Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
		// 	},
		// 	args: args{
		// 		w: w2,
		// 		r: r2,
		// 	},
		// 	resCode: 302,
		// 	w:       w2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Delegate:       tt.fields.Delegate,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				SiteName:       tt.fields.SiteName,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
				SiteURL:        tt.fields.SiteURL,
			}
			h.GetPrivacyPolicy(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}
