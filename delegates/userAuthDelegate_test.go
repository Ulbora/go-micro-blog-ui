package delegates

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
)

func TestMCDelegate_AddUserAuth(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":15}`)),
	}
	proxy.MockRespCode = 200

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		proxy       px.Proxy
		Log         lg.Log
		RestURL     string
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		a *UserAuth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				proxy:       proxy.New(),
				Log:         log,
				RestURL:     "http://localhost:3000",
				APIAdminKey: "54211789991515",
				APIKey:      "557444414141",
			},
			args: args{
				a: &UserAuth{
					AuthType: "truthSocial",
					UserID:   12,
					Entered:  time.Now(),
				},
			},
			want: &ResponseID{
				ID:      15,
				Success: true,
				Code:    200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MCDelegate{
				proxy:       tt.fields.proxy,
				Log:         tt.fields.Log,
				RestURL:     tt.fields.RestURL,
				APIKey:      tt.fields.APIKey,
				APIAdminKey: tt.fields.APIAdminKey,
			}
			if got := d.AddUserAuth(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.AddUserAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetUserAuthList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "type": "linkedIn", "userId": 9, "entered":"2023-03-02T15:25:32Z"},
		{"id":6, "type": "linkedIn", "userId": 9, "entered":"2023-03-02T15:25:32Z"}]`)),
	}
	proxy.MockRespCode = 200

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		proxy       px.Proxy
		Log         lg.Log
		RestURL     string
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		userID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]UserAuth
		wantLen int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				proxy:       proxy.New(),
				Log:         log,
				RestURL:     "http://localhost:3000",
				APIAdminKey: "54211789991515",
				APIKey:      "557444414141",
			},
			args: args{
				userID: 12,
				start:  0,
				end:    100,
			},
			wantLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MCDelegate{
				proxy:       tt.fields.proxy,
				Log:         tt.fields.Log,
				RestURL:     tt.fields.RestURL,
				APIKey:      tt.fields.APIKey,
				APIAdminKey: tt.fields.APIAdminKey,
			}
			//dd:= d.New()
			if got := d.GetUserAuthList(tt.args.userID, tt.args.start, tt.args.end); len(*got) != tt.wantLen {
				t.Errorf("MCDelegate.GetUserAuthList() = %v, want %v", got, tt.want)
			}
		})
	}
}
