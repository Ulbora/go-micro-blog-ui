package signins

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
)

func TestLinkedInSignin_Authorization(t *testing.T) {
	// var p px.GoProxy
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	p.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`<html`)),
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/loginUser/linkedIn", nil)
	w := httptest.NewRecorder()

	type fields struct {
		ClientID     string
		ClientSecret string
		AuthURL      string
		TokenURL     string
		RedirectURI  string
		proxy        px.Proxy
		Log          lg.Log
	}
	tests := []struct {
		name    string
		fields  fields
		resCode int
		w       *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				ClientID:     "12345",
				ClientSecret: "12345",
				//proxy:        &p,
				Log: log,
			},
			resCode: 302,
			w:       w,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LinkedInSignin{
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				AuthURL:      tt.fields.AuthURL,
				TokenURL:     tt.fields.TokenURL,
				RedirectURI:  tt.fields.RedirectURI,
				proxy:        tt.fields.proxy,
				Log:          tt.fields.Log,
			}
			ss := s.New()
			ss.SetProxy(&p)
			ss.Authorization(w, r)
			if tt.resCode != tt.w.Code{
				t.Fail()
			}
			// if res.Code != 0 {
			// 	t.Fail()
			// }
		})
	}
}

func TestLinkedInSignin_AccessToken(t *testing.T) {
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

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		ClientID     string
		ClientSecret string
		AuthURL      string
		TokenURL     string
		RedirectURI  string
		proxy        px.Proxy
		Log          lg.Log
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenResponse
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				ClientID:     "12345",
				ClientSecret: "12345",
				//proxy:        &p,
				Log: log,
			},
			args: args{
				code: "AQQ6uLeyJLnPqHHctWoDicrjyRcVW8qMzjfuczZyh-vfm7i_luzBLthyQSuruP68yL8r7EaGVHARmxIKGvuLMAplSAL9_J3zb8_Y7abyqdjvrjEoNSXaA6s_gg47_I5UK2DKmwsZjLx_4nsnIeKNi0ENP2OnQQRpBZJgS_IAka_Rv7IxX21LxxjvuS-77kxDh_aynCHzBfWAAOSZdzc",
			},
			want: &TokenResponse{
				AccessToken: "AQXjDdPrhpD-kyXXOO5T_i2eyANn8NQIaFLlqc1fyCCOFiTGK36I_jBrvzClHeReNBPPTHJ_bGfC31nZGkzpd93_TfjEurhLeO3hpFQq-klXqr0RbPsDWIBQ77HlwaKx_62Bj1aBvsEciS493bd3DTUMSrgOPmWHaGKHRYecH4Lf-lpgWo35LS6ssFmvtljG1NwWuTaXYH1kSr6SxqUZKbp5eQkS3E0WpWkxc9ofEHyHrwIb0vHPbSOUrNr0huIkCoON3hrMPUhoWcO3bO_uP0mXRnzGHWwGG38Dh6ploXTBYIeDrq_6SKqExcS6ua_ITuRHKNtAYLunZlqKZefqXEYDyopE6Q",
				ExpiresIn:   5183999,
				Scope:       "email,profile,w_member_social",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LinkedInSignin{
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				AuthURL:      tt.fields.AuthURL,
				TokenURL:     tt.fields.TokenURL,
				RedirectURI:  tt.fields.RedirectURI,
				proxy:        tt.fields.proxy,
				Log:          tt.fields.Log,
			}
			ss := s.New()
			ss.SetProxy(&p)
			if got := ss.AccessToken(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedInSignin.AccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
