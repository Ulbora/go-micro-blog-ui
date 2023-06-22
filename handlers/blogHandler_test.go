package handlers

// import (
// 	"net/http"
// 	"testing"
// 	"text/template"

// 	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
// )

import (
	"html/template"
	"net/http"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
	m "github.com/Ulbora/go-micro-blog-ui/managers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_GetBlogList(t *testing.T) {
	type fields struct {
		Log            lg.Log
		Manager        m.Manager
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
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				Manager:        tt.fields.Manager,
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
		})
	}
}
