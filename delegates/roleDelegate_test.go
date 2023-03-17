package delegates

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
)

func TestMCDelegate_AddRole(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":6}`)),
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
		r *Role
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
				r: &Role{
					Name: "testRole3",
				},
			},
			want: &ResponseID{
				ID:      6,
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
			if got := d.AddRole(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.AddRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetRole(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"id":9, "name":"testRole3"}`)),
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
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Role
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
				name: "testRole3",
			},
			want: &Role{
				ID:   9,
				Name: "testRole3",
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
			if got := d.GetRole(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetRoleList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id":17, "name":"test1234"}, {"id":18, "name":"test12344"}]`)),
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
	tests := []struct {
		name    string
		fields  fields
		want    *[]Role
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
			if got := d.GetRoleList(); len(*got) != tt.wantLen {
				t.Errorf("MCDelegate.GetRoleList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_DeleteRole(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`)),
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
		rid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
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
				rid: 10,
			},
			want: &Response{
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
			if got := d.DeleteRole(tt.args.rid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.DeleteRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
