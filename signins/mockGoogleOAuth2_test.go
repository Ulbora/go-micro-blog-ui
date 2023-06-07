package signins

import (
	"reflect"
	"testing"
)

func TestMockGoogleSignin_New(t *testing.T) {
	type fields struct {
		MockToken             string
		MockUserEmailVerified bool
		MockUserEmail         string
		MockUserFirstName     string
		MockUserLastName      string
		MockUserPictureURL    string
		MockUserPictureImage  []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   Signin
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MockGoogleSignin{
				MockToken:             tt.fields.MockToken,
				MockUserEmailVerified: tt.fields.MockUserEmailVerified,
				MockUserEmail:         tt.fields.MockUserEmail,
				MockUserFirstName:     tt.fields.MockUserFirstName,
				MockUserLastName:      tt.fields.MockUserLastName,
				MockUserPictureURL:    tt.fields.MockUserPictureURL,
				MockUserPictureImage:  tt.fields.MockUserPictureImage,
			}
			if got := s.New(); got == nil {
				t.Errorf("MockGoogleSignin.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockGoogleSignin_AccessToken(t *testing.T) {
	type fields struct {
		MockToken             string
		MockUserEmailVerified bool
		MockUserEmail         string
		MockUserFirstName     string
		MockUserLastName      string
		MockUserPictureURL    string
		MockUserPictureImage  []byte
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
				MockToken: "12345",
			},
			want: &TokenResponse{
				AccessToken: "12345",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MockGoogleSignin{
				MockToken:             tt.fields.MockToken,
				MockUserEmailVerified: tt.fields.MockUserEmailVerified,
				MockUserEmail:         tt.fields.MockUserEmail,
				MockUserFirstName:     tt.fields.MockUserFirstName,
				MockUserLastName:      tt.fields.MockUserLastName,
				MockUserPictureURL:    tt.fields.MockUserPictureURL,
				MockUserPictureImage:  tt.fields.MockUserPictureImage,
			}
			if got := s.AccessToken(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockGoogleSignin.AccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockGoogleSignin_GetUserInfo(t *testing.T) {
	type fields struct {
		MockToken             string
		MockUserEmailVerified bool
		MockUserEmail         string
		MockUserFirstName     string
		MockUserLastName      string
		MockUserPictureURL    string
		MockUserPictureImage  []byte
	}
	type args struct {
		tk string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockUserEmail: "test@test.com",
			},
			want: &GoogleUserInfoResponse{
				Email: "test@test.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MockGoogleSignin{
				MockToken:             tt.fields.MockToken,
				MockUserEmailVerified: tt.fields.MockUserEmailVerified,
				MockUserEmail:         tt.fields.MockUserEmail,
				MockUserFirstName:     tt.fields.MockUserFirstName,
				MockUserLastName:      tt.fields.MockUserLastName,
				MockUserPictureURL:    tt.fields.MockUserPictureURL,
				MockUserPictureImage:  tt.fields.MockUserPictureImage,
			}
			if got := s.GetUserInfo(tt.args.tk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockGoogleSignin.GetUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockGoogleSignin_GetUserPicture(t *testing.T) {
	type fields struct {
		MockToken             string
		MockUserEmailVerified bool
		MockUserEmail         string
		MockUserFirstName     string
		MockUserLastName      string
		MockUserPictureURL    string
		MockUserPictureImage  []byte
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MockGoogleSignin{
				MockToken:             tt.fields.MockToken,
				MockUserEmailVerified: tt.fields.MockUserEmailVerified,
				MockUserEmail:         tt.fields.MockUserEmail,
				MockUserFirstName:     tt.fields.MockUserFirstName,
				MockUserLastName:      tt.fields.MockUserLastName,
				MockUserPictureURL:    tt.fields.MockUserPictureURL,
				MockUserPictureImage:  tt.fields.MockUserPictureImage,
			}
			if got := s.GetUserPicture(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockGoogleSignin.GetUserPicture() = %v, want %v", got, tt.want)
			}
		})
	}
}
