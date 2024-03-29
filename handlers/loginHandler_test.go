package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	mux "github.com/GolangToolKits/grrt"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_LoginUserPage(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	sessionManager2, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("GET", "/auth/linkedin/callback?code=AQTe-q0dCREmAyHM1SH3DRlwJqcBdnuBieH7OoM5NbnAgGC04apOpPGmzO4rGfLybxxRadelqSR7mDe1hrPEw6Zkr-mZphb93o5hVz8dEtrjLFfnNX8ohyV-mrRTehgwDSvmoC_rulnTcCfkTVkMAHuzuCcJP2DCTnfUaeO0WpHf1_QBm3RmgCHN8BFyihvmSGJ5u4dHMCrfhT9iTvw&state=ieridf7fsf6dfs6", nil)
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/auth/linkedin/callback?code=AQTe-q0dCREmAyHM1SH3DRlwJqcBdnuBieH7OoM5NbnAgGC04apOpPGmzO4rGfLybxxRadelqSR7mDe1hrPEw6Zkr-mZphb93o5hVz8dEtrjLFfnNX8ohyV-mrRTehgwDSvmoC_rulnTcCfkTVkMAHuzuCcJP2DCTnfUaeO0WpHf1_QBm3RmgCHN8BFyihvmSGJ5u4dHMCrfhT9iTvw&state=ieridf7fsf6dfs6_", nil)
	w2 := httptest.NewRecorder()

	ses2 := sessionManager2.NewSession(r2, "go-micro-blog-ui")
	ses2.Set("loggedIn", true)
	ses2.Save(w2)

	//do the job of the browser and save response cookies to the request
	cook := w2.Result().Cookies()
	if len(cook) > 0 {
		r2.AddCookie(cook[0])
	}
	//end copy cookie from response to request that should be done by the browser
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	// var p px.GoProxy
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	p.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"access_token": "AQXjDdPrhpD-kyXXOO5T_i2eyANn8NQIaFLlqc1fyCCOFiTGK36I_jBrvzClHeReNBPPTHJ_bGfC31nZGkzpd93_TfjEurhLeO3hpFQq-klXqr0RbPsDWIBQ77HlwaKx_62Bj1aBvsEciS493bd3DTUMSrgOPmWHaGKHRYecH4Lf-lpgWo35LS6ssFmvtljG1NwWuTaXYH1kSr6SxqUZKbp5eQkS3E0WpWkxc9ofEHyHrwIb0vHPbSOUrNr0huIkCoON3hrMPUhoWcO3bO_uP0mXRnzGHWwGG38Dh6ploXTBYIeDrq_6SKqExcS6ua_ITuRHKNtAYLunZlqKZefqXEYDyopE6Q",
			"expires_in": 5183999,
			"scope": "email,profile,w_member_social"
		}`)),
	}

	var si s.LinkedInSignin
	si.ClientID = "12354"
	si.ClientSecret = "12345"
	si.Log = log

	sii := si.New()
	sii.SetProxy(&p)

	var signinMap = make(map[string]s.Signin)
	signinMap["linkedIn"] = sii

	type fields struct {
		Log            lg.Log
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		AdminTemplates *template.Template
		Templates      *template.Template
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
				Signins:        signinMap,
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
				Signins:        signinMap,
				SessionManager: sessionManager2,
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
			hh := &MCHandler{
				Log:            tt.fields.Log,
				APIKey:         tt.fields.APIKey,
				APIAdminKey:    tt.fields.APIAdminKey,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				AdminTemplates: tt.fields.AdminTemplates,
				Templates:      tt.fields.Templates,
			}
			h := hh.New()
			h.LoginUserPage(tt.args.w, tt.args.r)
			if tt.resCode != tt.w.Code {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_LoginUser(t *testing.T) {
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	// var p px.GoProxy
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	p.MockRespCode = 200
	p.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"access_token": "AQXjDdPrhpD-kyXXOO5T_i2eyANn8NQIaFLlqc1fyCCOFiTGK36I_jBrvzClHeReNBPPTHJ_bGfC31nZGkzpd93_TfjEurhLeO3hpFQq-klXqr0RbPsDWIBQ77HlwaKx_62Bj1aBvsEciS493bd3DTUMSrgOPmWHaGKHRYecH4Lf-lpgWo35LS6ssFmvtljG1NwWuTaXYH1kSr6SxqUZKbp5eQkS3E0WpWkxc9ofEHyHrwIb0vHPbSOUrNr0huIkCoON3hrMPUhoWcO3bO_uP0mXRnzGHWwGG38Dh6ploXTBYIeDrq_6SKqExcS6ua_ITuRHKNtAYLunZlqKZefqXEYDyopE6Q",
			"expires_in": 5183999,
			"scope": "email,profile,w_member_social"
		}`)),
	}

	var si s.LinkedInSignin
	si.ClientID = "12354"
	si.ClientSecret = "12345"
	si.Log = log

	sii := si.New()
	sii.SetProxy(&p)

	var signinMap = make(map[string]s.Signin)
	signinMap["linkedIn"] = sii

	r, _ := http.NewRequest("GET", "/loginUser/linkedIn", nil)
	vars := map[string]string{
		"signingSystem": "linkedIn",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	var p2 px.MockGoProxy
	p2.MockDoSuccess1 = true
	p2.MockRespCode = 0
	p2.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
			"access_token": "AQXjDdPrhpD-kyXXOO5T_i2eyANn8NQIaFLlqc1fyCCOFiTGK36I_jBrvzClHeReNBPPTHJ_bGfC31nZGkzpd93_TfjEurhLeO3hpFQq-klXqr0RbPsDWIBQ77HlwaKx_62Bj1aBvsEciS493bd3DTUMSrgOPmWHaGKHRYecH4Lf-lpgWo35LS6ssFmvtljG1NwWuTaXYH1kSr6SxqUZKbp5eQkS3E0WpWkxc9ofEHyHrwIb0vHPbSOUrNr0huIkCoON3hrMPUhoWcO3bO_uP0mXRnzGHWwGG38Dh6ploXTBYIeDrq_6SKqExcS6ua_ITuRHKNtAYLunZlqKZefqXEYDyopE6Q",
			"expires_in": 5183999,
			"scope": "email,profile,w_member_social"
		}`)),
	}

	var si2 s.GoogleSignin
	si2.ClientID = "12354"
	si2.ClientSecret = "12345"
	si2.Log = log

	sii2 := si2.New()
	sii2.SetProxy(&p2)

	var signinMap2 = make(map[string]s.Signin)
	signinMap2["googleOAuth2"] = sii2

	r2, _ := http.NewRequest("GET", "/loginUser/linkedIn", nil)
	vars2 := map[string]string{
		"signingSystem": "googleOAuth2",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

	type fields struct {
		Log            lg.Log
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
				Log:     log,
				Signins: signinMap,
				//SessionManager: sessionManager,
				Templates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				r: r,
				w: w,
			},
			resCode: 302,
			w:       w,
		},
		{
			name: "test 2",
			fields: fields{
				Log:     log,
				Signins: signinMap2,
				//SessionManager: sessionManager,
				Templates: template.Must(template.ParseFiles("testHtmls/test.html")),
			},
			args: args{
				r: r2,
				w: w2,
			},
			resCode: 302,
			w:       w2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hh := &MCHandler{
				Log:            tt.fields.Log,
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
			h := hh.New()
			h.LoginUser(tt.args.w, tt.args.r)
			if tt.w.Code != tt.resCode {
				t.Fail()
			}
		})
	}
}
