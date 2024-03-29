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

func TestMCHandler_GetBlogList(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/", nil)
	w2 := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Set("isAdmin", true)
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate
	var b1 mcd.Blog
	b1.Active = true
	b1.Content = "dGVzdCAx"
	b1.Name = "test 1"
	b1.UserID = 21

	var b2 mcd.Blog
	b2.Active = true
	b2.Content = "dGVzdCAy"
	b2.Name = "test 2"
	b2.UserID = 22

	var b3 mcd.Blog
	b3.Active = true
	b3.Content = "dGVzdCAz"
	b3.Name = "test 3"
	b3.UserID = 23

	var blo []mcd.Blog
	blo = append(blo, b1)
	blo = append(blo, b2)
	blo = append(blo, b3)
	mcdel.MockBlogList = &blo

	var c1 mcd.Comment
	c1.Active = true
	c1.BlogID = 1
	c1.Text = "test c 1"
	c1.UserID = 11

	var c2 mcd.Comment
	c2.Active = true
	c2.BlogID = 2
	c2.Text = "test c 2"
	c2.UserID = 22
	var cml []mcd.Comment
	cml = append(cml, c1)
	cml = append(cml, c2)
	mcdel.MockCommentList = &cml

	var l1 mcd.Like
	l1.BlogID = 44
	l1.UserID = 22

	var l2 mcd.Like
	l2.BlogID = 44
	l2.UserID = 23
	var lkl []mcd.Like
	lkl = append(lkl, l1)
	lkl = append(lkl, l2)
	mcdel.MockLikeList = &lkl

	//mcdel.GetUserActive = true
	//mcdel.GetUserEmail = "test@test.com"

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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.GetBlogList(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_AddBlogPage(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("GET", "/", nil)
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.AddBlogPage(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_AddBlog(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("POST", "/addBlog", strings.NewReader("name=test&userEmail=bob@b.com&content=this is a test blog"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("POST", "/addBlog", strings.NewReader("name=test&userEmail=bob@b.com&content=this is a test blog"))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	mcdel.MockAddBlogRes = &mcd.ResponseID{
		ID:      1,
		Success: true,
	}

	var mu1 mcd.User
	mu1.Active = true
	mu1.Email = "test@test.com"
	//mcdel2.GetUserActive = true
	//mcdel2.GetUserEmail = "test@test.com"
	mcdel.MockUser = &mu1

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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.AddBlog(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetBlog(t *testing.T) {

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
		"bid": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("GET", "/", nil)
	vars2 := map[string]string{
		"bid": "5",
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

	var b1 mcd.Blog
	b1.Active = true
	b1.Content = "dGVzdCAx"
	b1.Name = "test 1"
	b1.UserID = 21

	mcdel.MockBlog = &b1

	var c1 mcd.Comment
	c1.Active = true
	c1.BlogID = 1
	c1.Text = "test c 1"
	c1.UserID = 11

	var c2 mcd.Comment
	c2.Active = true
	c2.BlogID = 2
	c2.Text = "test c 2"
	c2.UserID = 22
	var cml []mcd.Comment
	cml = append(cml, c1)
	cml = append(cml, c2)
	mcdel.MockCommentList = &cml

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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.GetBlog(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_UpdateBlogPage(t *testing.T) {

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
		"bid": "5",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("GET", "/", nil)
	vars2 := map[string]string{
		"bid": "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", true)
	ses2.Set("userEmail", "bob@bob2.com")
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	r3, _ := http.NewRequest("GET", "/", nil)
	vars3 := map[string]string{
		"bid": "5",
	}
	r3 = mux.SetURLVars(r3, vars3)
	w3 := httptest.NewRecorder()

	ses3 := sessionManager.NewSession(r3, "go-micro-blog-ui")
	ses3.Set("loggedIn", false)
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

	var b1 mcd.Blog
	b1.Active = true
	b1.Content = "dGVzdCAx"
	b1.Name = "test 1"
	b1.UserID = 21

	mcdel.MockBlog = &b1

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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			h.UpdateBlogPage(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_UpdateBlog(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("PUT", "/updateBlog", strings.NewReader("id=5&userId=8&name=test&userEmail=bob@b.com&content=this is a test blog"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	r2, _ := http.NewRequest("PUT", "/updateBlog", strings.NewReader("id=5&userId=8&name=test&userEmail=bob@b.com&content=this is a test blog"))
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	mcdel.MockUpdateBlogRes = &mcd.Response{
		Success: true,
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.UpdateBlog(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_SearchBlogList(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("POST", "/searchBlog", strings.NewReader("name=test"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/", nil)
	w2 := httptest.NewRecorder()

	ses := sessionManager.NewSession(r, "go-micro-blog-ui")
	ses.Set("loggedIn", true)
	ses.Set("userEmail", "bob@bob.com")
	ses.Set("isAdmin", true)
	ses.Save(w)

	cook := w.Result().Cookies()
	if len(cook) > 0 {
		r.AddCookie(cook[0])
	}

	ses2 := sessionManager.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", false)
	ses2.Save(w2)

	cook2 := w2.Result().Cookies()
	if len(cook2) > 0 {
		r2.AddCookie(cook2[0])
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var mcdel mcd.MockDelegate
	var b1 mcd.Blog
	b1.Active = true
	b1.Content = "dGVzdCAx"
	b1.Name = "test 1"
	b1.UserID = 21

	var b2 mcd.Blog
	b2.Active = true
	b2.Content = "dGVzdCAy"
	b2.Name = "test 2"
	b2.UserID = 22

	var b3 mcd.Blog
	b3.Active = true
	b3.Content = "dGVzdCAz"
	b3.Name = "test 3"
	b3.UserID = 23

	var blo []mcd.Blog
	blo = append(blo, b1)
	blo = append(blo, b2)
	blo = append(blo, b3)
	mcdel.MockBlogList = &blo

	var c1 mcd.Comment
	c1.Active = true
	c1.BlogID = 1
	c1.Text = "test c 1"
	c1.UserID = 11

	var c2 mcd.Comment
	c2.Active = true
	c2.BlogID = 2
	c2.Text = "test c 2"
	c2.UserID = 22
	var cml []mcd.Comment
	cml = append(cml, c1)
	cml = append(cml, c2)
	mcdel.MockCommentList = &cml

	var l1 mcd.Like
	l1.BlogID = 44
	l1.UserID = 22

	var l2 mcd.Like
	l2.BlogID = 44
	l2.UserID = 23
	var lkl []mcd.Like
	lkl = append(lkl, l1)
	lkl = append(lkl, l2)
	mcdel.MockLikeList = &lkl

	//mcdel.GetUserActive = true
	//mcdel.GetUserEmail = "test@test.com"

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
		{
			name: "test 2",
			fields: fields{
				Log:            log,
				Delegate:       &mcdel,
				SessionManager: sessionManager,
				Templates:      template.Must(template.ParseFiles("testHtmls/test.html")),
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
			h.SearchBlogList(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}
