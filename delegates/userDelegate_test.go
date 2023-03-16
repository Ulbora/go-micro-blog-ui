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

func TestMCDelegate_AddUser(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":5}`)),
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
		u *User
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
				u: &User{
					Email:     "test1234@test.com",
					FirstName: "test",
					LastName:  "test123",
					RoleID:    1,
				},
			},
			want: &ResponseID{
				ID:      5,
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
			if got := d.AddUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_UpdateUser(t *testing.T) {

	//var proxy px.GoProxy

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
		u *User
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
				u: &User{
					ID:        17,
					Email:     "test1234@test.com",
					FirstName: "test",
					LastName:  "test1231231235",
					RoleID:    1,
				},
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
			if got := d.UpdateUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetUser(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"id":17, "email":"test1234@test.com"}`)),
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
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User

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
				email: "test1234@test.com",
			},
			want: &User{
				ID:        17,
				Email:     "test1234@test.com",
				FirstName: "test",
				LastName:  "test1231231235",
				RoleID:    1,
				Active:    true,
				Password: "",
				Image: nil,
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
			if got := d.GetUser(tt.args.email); got.Email != tt.want.Email ||
			got.ID != tt.want.ID {
				t.Errorf("MCDelegate.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
