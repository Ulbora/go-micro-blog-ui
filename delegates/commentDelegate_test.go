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

func TestMCDelegate_AddComment(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":10}`)),
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
		c *Comment
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
				c: &Comment{
					BlogID: 9,
					UserID: 12,
					Text:   "this is a good blog",
				},
			},
			want: &ResponseID{
				ID:      10,
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
			if got := d.AddComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_UpdateComment(t *testing.T) {

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
		c *Comment
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
				c: &Comment{
					ID:     9,
					BlogID: 9,
					UserID: 12,
					Text:   "update blog comment for new",
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
			if got := d.UpdateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetCommentList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "blogId": 2, "userId": 9,  "text":"test blog 12345--------", "active": true},
		{"id":6, "blogId": 2, "userId": 9,  "text":"test blog 12345--------", "active": true}]`)),
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
		blogID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]Comment
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
				blogID: 4,
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
			if got := d.GetCommentList(tt.args.blogID, tt.args.start, tt.args.end); len(*got) != tt.wantLen {
				t.Errorf("MCDelegate.GetCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetAdminCommentList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "blogId": 2, "userId": 9,  "text":"test blog 12345--------", "active": true},
		{"id":6, "blogId": 2, "userId": 9,  "text":"test blog 12345--------", "active": true}]`)),
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
		blogID int64
		start  int64
		end    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]Comment
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
				blogID: 4,
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
			if got := d.GetAdminCommentList(tt.args.blogID, tt.args.start, tt.args.end); len(*got) != tt.wantLen {
				t.Errorf("MCDelegate.GetAdminCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_ActivateComment(t *testing.T) {

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
		c *Comment
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
				c: &Comment{
					ID:     10,
					BlogID: 9,
					UserID: 12,
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
			if got := d.ActivateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.ActivateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_DeActivateComment(t *testing.T) {

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
		c *Comment
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
				c: &Comment{
					ID:     10,
					BlogID: 9,
					UserID: 12,
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
			if got := d.DeActivateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.DeActivateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
