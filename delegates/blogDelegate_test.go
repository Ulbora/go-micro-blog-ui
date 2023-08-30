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

func TestMCDelegate_AddBlog(t *testing.T) {

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
		b *Blog
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
				b: &Blog{
					Name:    "test blog 12345",
					Content: "just some more content !!!!!!",
					UserID:  12,
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
			if got := d.AddBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.AddBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_UpdateBlog(t *testing.T) {

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
		b *Blog
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
				b: &Blog{
					ID:      5,
					Name:    "test blog 12345--------",
					Content: "just some more content !!!!!!------",
					UserID:  12,
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
			if got := d.UpdateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.UpdateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetBlog(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"id":5, "name":"test blog 12345--------"}`)),
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
		bid int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *Blog
		wantID   int64
		wantName string
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
				bid: 5,
			},
			wantID:   5,
			wantName: "test blog 12345--------",
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
			if got := d.GetBlog(tt.args.bid); got.ID != tt.wantID || got.Name != tt.wantName {
				t.Errorf("MCDelegate.GetBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetBlogByName(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "name":"test blog 12345--------"}]`)),
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
		name  string
		start int64
		end   int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *Blog
		wantID   int64
		wantName string
		wantLen  int
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
				name:  "test blog 12345--------",
				start: 0,
				end:   100,
			},
			wantID:   5,
			wantName: "test blog 12345--------",
			wantLen:  1,
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
			if got := d.GetBlogByName(tt.args.name, tt.args.start, tt.args.end); len(*got) != tt.wantLen ||
				(*got)[0].ID != tt.wantID || (*got)[0].Name != tt.wantName {
				t.Errorf("MCDelegate.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}



func TestMCDelegate_GetAdminBlogByName(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "name":"test blog 12345--------"}]`)),
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
		name  string
		start int64
		end   int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *Blog
		wantID   int64
		wantName string
		wantLen  int
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
				name:  "test blog 12345--------",
				start: 0,
				end:   100,
			},
			wantID:   5,
			wantName: "test blog 12345--------",
			wantLen:  1,
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
			if got := d.GetAdminBlogByName(tt.args.name, tt.args.start, tt.args.end); len(*got) != tt.wantLen ||
				(*got)[0].ID != tt.wantID || (*got)[0].Name != tt.wantName {
				t.Errorf("MCDelegate.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetBlogList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "name":"test blog 12345--------"}]`)),
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
		start int64
		end   int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *[]Blog
		wantID   int64
		wantName string
		wantLen  int
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
				start: 0,
				end:   100,
			},
			wantID:   5,
			wantName: "test blog 12345--------",
			// wantName: "test blog entry 222",
			wantLen: 1,
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
			if got := d.GetBlogList(tt.args.start, tt.args.end); len(*got) != tt.wantLen ||
				(*got)[0].ID != tt.wantID || (*got)[0].Name != tt.wantName {
				t.Errorf("MCDelegate.GetBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_GetAdminBlogList(t *testing.T) {

	// var proxy px.GoProxy

	var proxy px.MockGoProxy
	proxy.MockDoSuccess1 = true
	proxy.MockResp = &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id":5, "name":"test blog 12345--------"}]`)),
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
		start int64
		end   int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *[]Blog
		wantID   int64
		wantName string
		wantLen  int
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
				start: 0,
				end:   100,
			},
			wantID:   5,
			wantName: "test blog 12345--------",
			// wantName: "test blog entry 222",
			wantLen: 1,
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
			if got := d.GetAdminBlogList(tt.args.start, tt.args.end); len(*got) != tt.wantLen ||
				(*got)[0].ID != tt.wantID || (*got)[0].Name != tt.wantName {
				t.Errorf("MCDelegate.GetAdminBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_ActivateBlog(t *testing.T) {

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
		b *Blog
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
				b: &Blog{
					ID: 5,
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
			if got := d.ActivateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.ActivateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_DeActivateBlog(t *testing.T) {

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
		b *Blog
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
				b: &Blog{
					ID: 5,
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
			if got := d.DeActivateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.DeActivateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCDelegate_DeleteBlog(t *testing.T) {
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
		bid int64
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
				bid: 5,
			},
			want: &Response{
				Success: true,
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
			if got := d.DeleteBlog(tt.args.bid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MCDelegate.DeleteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}
