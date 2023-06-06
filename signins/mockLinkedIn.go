package signins

import (
	"net/http"

	px "github.com/GolangToolKits/go-http-proxy"
)

// MockLinkedInSignin MockLinkedInSignin
type MockLinkedInSignin struct {
	MockToken             string
	MockUserEmailVerified bool
	MockUserEmail         string
	MockUserFirstName     string
	MockUserLastName      string
	MockUserPictureURL    string
	MockUserPictureImage  []byte
}

// New New
func (s *MockLinkedInSignin) New() Signin {
	return s
}

// Authorization Authorization
func (s *MockLinkedInSignin) Authorization(w http.ResponseWriter, r *http.Request) {

}

// AccessToken AccessToken
func (s *MockLinkedInSignin) AccessToken(code string) *TokenResponse {
	var rtn TokenResponse
	rtn.AccessToken = s.MockToken

	return &rtn
}

// GetUserInfo GetUserInfo
func (s *MockLinkedInSignin) GetUserInfo(tk string) any {
	var rtn UserInfoResponse
	rtn.Email = s.MockUserEmail
	rtn.EmailVerified = s.MockUserEmailVerified
	rtn.FirstName = s.MockUserFirstName
	rtn.LastName = s.MockUserLastName
	rtn.PictureURL = s.MockUserPictureURL
	return &rtn
}

// GetUserPicture GetUserPicture
func (s *MockLinkedInSignin) GetUserPicture(url string) []byte {
	return s.MockUserPictureImage
}

// SetProxy SetProxy
func (s *MockLinkedInSignin) SetProxy(p px.Proxy) {

}
