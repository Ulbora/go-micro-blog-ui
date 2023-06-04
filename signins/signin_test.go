package signins

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"testing"
)

func Test_buildRequest(t *testing.T) {
	type Blog struct {
		Name   string
		UserID int64
	}
	var ba Blog
	ba.Name = "A test blog"
	ba.UserID = 12
	aJSON, err := json.Marshal(ba)
	if err != nil {
		fmt.Println("err: ", err)
	}
	type args struct {
		method string
		url    string
		aJSON  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				method: http.MethodPost,
				url:    "/test",
				aJSON: aJSON,
			},
		},
		{
			name: "test 2",
			args: args{
				method: http.MethodGet,
				url:    "/test",
				aJSON: aJSON,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildRequest(tt.args.method, tt.args.url, tt.args.aJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("buildRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
