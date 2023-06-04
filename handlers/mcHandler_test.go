package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	m "github.com/Ulbora/go-micro-blog-ui/managers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_processAPIKey(t *testing.T) {
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/auth/", nil)
	r.Header.Set("apiKey", "1234")

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				APIKey: "1234",
				Log:    log,
			},
			args: args{
				r: r,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				Manager:        tt.fields.Manager,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			if got := h.processAPIKey(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.processAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_processAPIAdminKey(t *testing.T) {
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/auth/", nil)
	r.Header.Set("apiAdminKey", "1234")

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:         log,
				APIAdminKey: "1234",
			},
			args: args{
				r: r,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				Manager:        tt.fields.Manager,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			if got := h.processAPIAdminKey(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.processAPIAdminKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_checkContent(t *testing.T) {
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/auth/", nil)
	r.Header.Set("Content-Type", "application/json")

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:         log,
				APIAdminKey: "1234",
			},
			args: args{
				r: r,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				Manager:        tt.fields.Manager,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			if got := h.checkContent(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.checkContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_processBody(t *testing.T) {
	type testObj struct {
		Valid bool
		Code  string
	}
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var robj testObj
	robj.Valid = true
	robj.Code = "3"

	aJSON, _ := json.Marshal(robj)

	r, _ := http.NewRequest("POST", "/auth/", bytes.NewBuffer(aJSON))

	r2, _ := http.NewRequest("POST", "/auth/", nil)
	//r.Header.Set("Content-Type", "application/json")

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
		Title          string
		Desc           string
		KeyWords       string
	}
	type args struct {
		r   *http.Request
		obj any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r,
				obj: &testObj{},
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r2,
				obj: &testObj{},
			},
			want: false,
			wantErr: true,
		},
		{
			name: "test 2",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r,
				//obj: &testObj{},
			},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				Log:            tt.fields.Log,
				Manager:        tt.fields.Manager,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				KeyWords:       tt.fields.KeyWords,
			}
			got, err := h.processBody(tt.args.r, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("MCHandler.processBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MCHandler.processBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
