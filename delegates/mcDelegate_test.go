package delegates

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
)

func TestMCDelegate_buildRequest(t *testing.T) {
	var proxy px.GoProxy

	//var proxy px.MockGoProxy

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	var ba Blog
	ba.Name = "A test blog"
	ba.UserID = 12
	aJSON, err := json.Marshal(ba)
	if err != nil {
		fmt.Println("err: ", err)
	}

	type fields struct {
		proxy       px.Proxy
		Log         lg.Log
		RestURL     string
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		method  string
		url     string
		aJSON   []byte
		keyType string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		want             *http.Request
		wantErr          bool
		contentType      string
		apiKeyValue      string
		apiAdminKeyValue string
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
				method:  http.MethodPost,
				url:     "/blog/add",
				aJSON:   aJSON,
				keyType: apiKey,
			},
			contentType:      "application/json",
			apiKeyValue:      "557444414141",
			apiAdminKeyValue: "",
		},
		{
			name: "test 2",
			fields: fields{
				proxy:       proxy.New(),
				Log:         log,
				RestURL:     "http://localhost:3000",
				APIAdminKey: "54211789991515",
				APIKey:      "557444414141",
			},
			args: args{
				method: http.MethodGet,
				url:    "/blog/get",
				//aJSON:   aJSON,
				keyType: adminKey,
			},
			//contentType:      "application/json",
			//apiKeyValue:      "557444414141",
			apiAdminKeyValue: "54211789991515",
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
			got, err := d.buildRequest(tt.args.method, tt.args.url, tt.args.aJSON, tt.args.keyType)
			if (err != nil) != tt.wantErr {
				t.Errorf("MCDelegate.buildRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || got.Header.Get("Content-Type") != tt.contentType ||
				got.Header.Get("apiKey") != tt.apiKeyValue ||
				got.Header.Get("apiAdminKey") != tt.apiAdminKeyValue {
				t.Errorf("MCDelegate.buildRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_New(t *testing.T) {

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
	tests := []struct {
		name   string
		fields fields
		want   Delegate
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log: log,
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
			if got := d.New(); got == nil {
				t.Errorf("MCDelegate.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_SetProxy(t *testing.T) {
	var proxy px.GoProxy

	type fields struct {
		proxy       px.Proxy
		Log         lg.Log
		RestURL     string
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		proxy px.Proxy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{

			},
			args: args{
				proxy: &proxy,
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
			d.SetProxy(tt.args.proxy)
		})
	}
}
