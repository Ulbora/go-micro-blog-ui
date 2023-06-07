package signins

import (
	"net/http"

	px "github.com/GolangToolKits/go-http-proxy"
)

// MockGoogleSignin MockGoogleSignin
type MockGoogleSignin struct {
	MockToken             string
	MockUserEmailVerified bool
	MockUserEmail         string
	MockUserFirstName     string
	MockUserLastName      string
	MockUserPictureURL    string
	MockUserPictureImage  []byte
}

// New New
func (s *MockGoogleSignin) New() Signin {
	return s
}

// Authorization Authorization
func (s *MockGoogleSignin) Authorization(w http.ResponseWriter, r *http.Request) {
}

// AccessToken AccessToken
func (s *MockGoogleSignin) AccessToken(code string) *TokenResponse {
	var rtn TokenResponse
	rtn.AccessToken = s.MockToken

	return &rtn
}

// GetUserInfo GetUserInfo
func (s *MockGoogleSignin) GetUserInfo(tk string) any {
	var rtn GoogleUserInfoResponse
	rtn.Email = s.MockUserEmail
	rtn.EmailVerified = s.MockUserEmailVerified
	rtn.FirstName = s.MockUserFirstName
	rtn.LastName = s.MockUserLastName
	rtn.PictureURL = s.MockUserPictureURL
	return &rtn
}

// GetUserPicture GetUserPicture
func (s *MockGoogleSignin) GetUserPicture(url string) []byte {
	return s.MockUserPictureImage
}

// SetProxy SetProxy
func (s *MockGoogleSignin) SetProxy(p px.Proxy) {
}
