package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	del "github.com/Ulbora/go-micro-blog-ui/delegates"
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

	r3, _ := http.NewRequest("GET", "/auth/linkedin/callback?code=AQTe-q0dCREmAyHM1SH3DRlwJqcBdnuBieH7OoM5NbnAgGC04apOpPGmzO4rGfLybxxRadelqSR7mDe1hrPEw6Zkr-mZphb93o5hVz8dEtrjLFfnNX8ohyV-mrRTehgwDSvmoC_rulnTcCfkTVkMAHuzuCcJP2DCTnfUaeO0WpHf1_QBm3RmgCHN8BFyihvmSGJ5u4dHMCrfhT9iTvw&state=ieridf7fsf6dfs6", nil)
	w3 := httptest.NewRecorder()

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

	var mcdel del.MockDelegate
	mcdel.AddUserID = 4
	mcdel.AddUserSuccess = true

	mcdel.GetRoleID = 3

	// mcdel.GetUserActive = true
	// mcdel.GetUserEmail = "test@test.com"

	var mcdel2 del.MockDelegate
	mcdel2.AddUserID = 4
	mcdel2.AddUserSuccess = true

	mcdel2.GetRoleID = 3

	mcdel2.GetUserActive = true
	mcdel2.GetUserEmail = "test@test.com"

	// var si s.LinkedInSignin
	// si.ClientID = "12354"
	// si.ClientSecret = "12345"
	// si.Log = log

	// sii := si.New()
	// sii.SetProxy(&p)

	// var signinMap = make(map[string]s.Signin)
	// signinMap["linkedIn"] = sii

	var si s.MockLinkedInSignin
	si.MockToken = "4567890"
	si.MockUserEmail = "test@test.com"
	si.MockUserEmailVerified = true
	si.MockUserFirstName = "test"
	si.MockUserLastName = "tester"
	si.MockUserPictureURL = "http://image"
	sii := si.New()

	var signinMap = make(map[string]s.Signin)
	signinMap["linkedIn"] = sii

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Signins        map[string]s.Signin
		SessionManager gss.SessionManager
		Delegate       del.Delegate
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
				Delegate:       &mcdel,
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
				Delegate:       &mcdel2,
			},
			args: args{
				w: w2,
				r: r2,
			},
			w:       w2,
			resCode: 302,
		},
		{
			name: "test 3",
			fields: fields{
				Log:            log,
				Signins:        signinMap,
				SessionManager: sessionManager,
				Delegate:       &mcdel2,
			},
			args: args{
				w: w3,
				r: r3,
			},
			w:       w3,
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
				Delegate:       tt.fields.Delegate,
			}
			hh := h.New()
			hh.LinkedInCallback(tt.args.w, tt.args.r)
			if tt.w.Code != tt.resCode {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GoogleSigninCallback(t *testing.T) {

	var cf gss.ConfigOptions
	cf.MaxAge = 3600
	cf.Path = "/"
	sessionManager, err := gss.NewSessionManager("dsdfsadfs61dsscfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
		log.Println("Session err: ", err)
	}

	r, _ := http.NewRequest("GET", "/signin-google/callback?state=ieridf7fsf6dfs65775&code=4%2F0AbUR2VOwN3hyOyG9K9TFJP2IXRELq4hIT1aYoJm1dyE5GB8M34zIq1k2s9qfTR2SLWk6Wg&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+openid&authuser=0&prompt=consent", nil)
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/signin-google/callback?state=ieridf7fsf6dfs65775_&code=4%2F0AbUR2VOwN3hyOyG9K9TFJP2IXRELq4hIT1aYoJm1dyE5GB8M34zIq1k2s9qfTR2SLWk6Wg&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+openid&authuser=0&prompt=consent", nil)
	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/signin-google/callback?state=ieridf7fsf6dfs65775&code=4%2F0AbUR2VOwN3hyOyG9K9TFJP2IXRELq4hIT1aYoJm1dyE5GB8M34zIq1k2s9qfTR2SLWk6Wg&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+openid&authuser=0&prompt=consent", nil)
	w3 := httptest.NewRecorder()

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
			"access_token": "ya29.a0AWY7CkkmSQMrk5NgpSfBjTuXNffXtxX1dEyDY1722VyTxYeAKJi4UtBGSd3IpPUbpl_TLBmHkUoVzwSDfjgtieMj9wLflxnrrUgqc4Sfhuo810HJRkYwxyaX75v395X9XwccJFprjzyZO03H3Y97LqNT5HtnaCgYKAQYSARASFQG1tDrpWPv7d-Kgb5DP84ohRR5qrw0163",
			"expires_in": 3599,
			"scope": "openid https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile",
			"token_type": "Bearer",
			"id_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg1YmE5MzEzZmQ3YTdkNGFmYTg0ODg0YWJjYzg0MDMwMDQzNjMxODAiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0MTg0NDQ4NTIwMTgtdW1wcnExaDg3b2l1MTZhYTg5MTBhdmhiaW81djRmMmouYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MTg0NDQ4NTIwMTgtdW1wcnExaDg3b2l1MTZhYTg5MTBhdmhiaW81djRmMmouYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDEwODY5OTc1OTIyNDMxOTI4NjIiLCJlbWFpbCI6Imtlbnd3aWxsaWFtc29uMTJAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJaaDlYZFRRZEdVTy1JZnY0NGdseWRBIiwibmFtZSI6IktlbiBXaWxsaWFtc29uIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FBY0hUdGZLTzkwdHNMUnhoajAwNWR2bExXaDhXalIyQ1ZHOUJUUHJIVXpTanc9czk2LWMiLCJnaXZlbl9uYW1lIjoiS2VuIiwiZmFtaWx5X25hbWUiOiJXaWxsaWFtc29uIiwibG9jYWxlIjoiZW4iLCJpYXQiOjE2ODYxNjM1ODQsImV4cCI6MTY4NjE2NzE4NH0.B2Wrf-SEZZ_IpLUeH6QqNZOgERIQ7CUQi_thVB-BQ9kVmIO33JYr-8B3oZvKvnxqveA6lEJRzE-HzfaRg-tKwWTQYsMPQS6-MdNCWJ6FDREZxzSYnBbZTiZxKQvRoUbKxVZsIFkmk7MDkiJfU7zZpj1MekJATomLQS3qgzoRRS_YneuNO8yiyjXSO24l2smcmFjCeeo_rBwDaDeEYXrVDH0nWO0fKsaWsRS9eHuRncpAIBWOVNOfWlqrW4LhU_0smMDgwWIgopV6-qaC5aKwPe0J54cCCnnZvZnVzMdvF5es0qLLreqzVZaMEFauuJkv00TkDyEfUCBd4ObOFCoosg"
		}`)),
	}

	var mcdel del.MockDelegate
	mcdel.AddUserID = 4
	mcdel.AddUserSuccess = true

	mcdel.GetRoleID = 3

	// mcdel.GetUserActive = true
	// mcdel.GetUserEmail = "test@test.com"

	var mcdel2 del.MockDelegate
	mcdel2.AddUserID = 4
	mcdel2.AddUserSuccess = true

	mcdel2.GetRoleID = 3

	mcdel2.GetUserActive = true
	mcdel2.GetUserEmail = "test@test.com"

	// var si s.LinkedInSignin
	// si.ClientID = "12354"
	// si.ClientSecret = "12345"
	// si.Log = log

	// sii := si.New()
	// sii.SetProxy(&p)

	// var signinMap = make(map[string]s.Signin)
	// signinMap["linkedIn"] = sii

	var si s.MockGoogleSignin
	si.MockToken = "4567890"
	si.MockUserEmail = "test@test.com"
	si.MockUserEmailVerified = true
	si.MockUserFirstName = "test"
	si.MockUserLastName = "tester"
	si.MockUserPictureURL = "http://image"
	sii := si.New()

	var signinMap = make(map[string]s.Signin)
	signinMap["googleOAuth2"] = sii

	type fields struct {
		Log            lg.Log
		Manager        m.Manager
		APIKey         string
		APIAdminKey    string
		Delegate       del.Delegate
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
				Signins:        signinMap,
				SessionManager: sessionManager,
				Delegate:       &mcdel,
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
				Delegate:       &mcdel2,
			},
			args: args{
				w: w2,
				r: r2,
			},
			w:       w2,
			resCode: 302,
		},
		{
			name: "test 3",
			fields: fields{
				Log:            log,
				Signins:        signinMap,
				SessionManager: sessionManager,
				Delegate:       &mcdel2,
			},
			args: args{
				w: w3,
				r: r3,
			},
			w:       w3,
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
				Delegate:       tt.fields.Delegate,
				Signins:        tt.fields.Signins,
				SessionManager: tt.fields.SessionManager,
				//AdminTemplates: tt.fields.AdminTemplates,
				//Templates:      tt.fields.Templates,
				Title:    tt.fields.Title,
				Desc:     tt.fields.Desc,
				KeyWords: tt.fields.KeyWords,
			}
			h.GoogleSigninCallback(tt.args.w, tt.args.r)
		})
	}
}
