package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_GetUserByIDPage(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var mus mcd.User
	mus.Active = true
	mus.Email = "bob@bob.com"
	mcdel.MockUser = &mus

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.GetUserByIDPage(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetUnactivatedUserList(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var ulst []mcd.User
	var u1 mcd.User
	u1.Active = false
	u1.Email = "bob@bob.com"
	ulst = append(ulst, u1)

	var u2 mcd.User
	u2.Active = false
	u2.Email = "bob2@bob.com"
	ulst = append(ulst, u2)

	mcdel.MockUnactivatedUserList = &ulst

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.GetUnactivatedUserList(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetBannedUserList(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var ulst []mcd.User
	var u1 mcd.User
	u1.Active = false
	u1.Email = "bob@bob.com"
	ulst = append(ulst, u1)

	var u2 mcd.User
	u2.Active = false
	u2.Email = "bob2@bob.com"
	ulst = append(ulst, u2)

	mcdel.MockBannedUserList = &ulst

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.GetBannedUserList(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_EnableUser(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var mus mcd.User
	mus.Active = true
	mus.Email = "bob@bob.com"
	mcdel.MockUser = &mus

	var meu = mcd.Response{
		Success: false,
		Code:    200,
	}
	mcdel.MockEnableUser = &meu

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.EnableUser(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_DisableUser(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var mus mcd.User
	mus.Active = true
	mus.Email = "bob@bob.com"
	mcdel.MockUser = &mus

	var meu = mcd.Response{
		Success: false,
		Code:    200,
	}
	mcdel.MockDisableUser = &meu

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.DisableUser(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_DisableUserForCause(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var mus mcd.User
	mus.Active = true
	mus.Email = "bob@bob.com"
	mcdel.MockUser = &mus

	var meu = mcd.Response{
		Success: false,
		Code:    200,
	}
	mcdel.MockDisableUserForCause = &meu

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.DisableUserForCause(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_ReinstateBannedUser(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

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
	vars2 := map[string]string{
		"uid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Set("userEmail", "bob@bob.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate

	var mus mcd.User
	mus.Active = true
	mus.Email = "bob@bob.com"
	mcdel.MockUser = &mus

	var meu = mcd.Response{
		Success: false,
		Code:    200,
	}
	mcdel.MockReactivateBannedUser = &meu

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Delegate       mcd.Delegate
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.ReinstateBannedUser(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}
