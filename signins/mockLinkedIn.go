package signins

import (
	"net/http"

	px "github.com/GolangToolKits/go-http-proxy"
)

/*
 Copyright (C) 2023 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2023 Ken Williamson
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
