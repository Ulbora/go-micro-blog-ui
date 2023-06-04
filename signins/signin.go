package signins

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

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

// Signin Signin
type Signin interface {
	//Authorization() *Response
	//Authorization2(w http.ResponseWriter, r *http.Request) *Response
	Authorization(w http.ResponseWriter, r *http.Request)
	AccessToken(code string) *TokenResponse
	SetProxy(p px.Proxy)
	//Callback() *TokenParams
}

// AuthParams AuthParams
type AuthParams struct {
	ResponseType string
	ClientID     string
	RedirectURI  string
	State        string
	Scope        string
}

// TokenParams TokenParams
type TokenParams struct {
	GrantType    string
	AuthCode     string
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// TokenResponse TokenResponse
type TokenResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string
	RefreshTokenExpiresIn int
	Scope                 string `json:"scope"`
}

// Response Response
type Response struct {
	Code int `json:"code"`
}

// buildRequest buildRequest
func buildRequest(method string, url string, aJSON []byte) (*http.Request, error) {
	var req *http.Request
	var err error
	if method == http.MethodPost || method == http.MethodPut {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(aJSON))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	return req, err
}

func buildFormRequest(method string, url string, data url.Values) (*http.Request, error) {
	var req *http.Request
	var err error
	if method == http.MethodPost || method == http.MethodPut {
		req, err = http.NewRequest(method, url, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	// else {
	// 	req, err = http.NewRequest(method, url, nil)
	// }
	return req, err
}
