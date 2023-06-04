package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	m "github.com/Ulbora/go-micro-blog-ui/managers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

func TestMCHandler_LinkedInCallback(t *testing.T) {
	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("GET", "/auth/linkedin/callback?code=AQTe-q0dCREmAyHM1SH3DRlwJqcBdnuBieH7OoM5NbnAgGC04apOpPGmzO4rGfLybxxRadelqSR7mDe1hrPEw6Zkr-mZphb93o5hVz8dEtrjLFfnNX8ohyV-mrRTehgwDSvmoC_rulnTcCfkTVkMAHuzuCcJP2DCTnfUaeO0WpHf1_QBm3RmgCHN8BFyihvmSGJ5u4dHMCrfhT9iTvw&state=ieridf7fsf6dfs6", nil)
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/auth/linkedin/callback?code=AQTe-q0dCREmAyHM1SH3DRlwJqcBdnuBieH7OoM5NbnAgGC04apOpPGmzO4rGfLybxxRadelqSR7mDe1hrPEw6Zkr-mZphb93o5hVz8dEtrjLFfnNX8ohyV-mrRTehgwDSvmoC_rulnTcCfkTVkMAHuzuCcJP2DCTnfUaeO0WpHf1_QBm3RmgCHN8BFyihvmSGJ5u4dHMCrfhT9iTvw&state=ieridf7fsf6dfs6_", nil)
	w2 := httptest.NewRecorder()
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
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
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
			},
			args: args{
				w: w,
				r: r,
			},
			w:       w,
			resCode: 302,
		},
		{
			name: "test 2",
			fields: fields{
				Log:            log,
				Signins:        signinMap,
				SessionManager: sessionManager,
			},
			args: args{
				w: w2,
				r: r2,
			},
			w:       w2,
			resCode: 302,
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
			}
			hh := h.New()
			hh.LinkedInCallback(tt.args.w, tt.args.r)
			if tt.w.Code != tt.resCode {
				t.Fail()
			}
		})
	}
}