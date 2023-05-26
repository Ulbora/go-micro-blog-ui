package signins

import (
	"net/http"

	px "github.com/GolangToolKits/go-http-proxy"
	lg "github.com/GolangToolKits/go-level-logger"
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

const (
	linkedInDefaultAuthURL  = "https://www.linkedin.com/oauth/v2/authorization"
	linkedInDefaultTokenURL = "https://www.linkedin.com/oauth/v2/accessToken"
	scope                   = "r_liteprofile%20r_emailaddress%20w_member_social"
	state                   = "ieridf7fsf6dfs6"
)

// LinkedInSignin LinkedInSignin
type LinkedInSignin struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	proxy        px.Proxy
	Log          lg.Log
}

// New NewLinkedIn
func (s *LinkedInSignin) New() Signin {
	if s.AuthURL == "" {
		s.AuthURL = linkedInDefaultAuthURL
	}
	if s.TokenURL == "" {
		s.TokenURL = linkedInDefaultTokenURL
	}
	var p px.GoProxy
	s.proxy = &p
	return s
}

// Authorization LinkedIn Authorization
func (s *LinkedInSignin) Authorization(redirectURI string) {
	var rtn Response
	var authURL = linkedInDefaultAuthURL + "?response_type=code&client_id=" +
		s.ClientID + "&redirect_uri=" + redirectURI + "&state=" + state +
		"&scope=" + scope
	s.Log.Info("url: ", authURL)
	brq, err := buildRequest(http.MethodGet, authURL, nil)
	if err == nil {
		gsuc, stat := s.proxy.Do(brq, &rtn)
		s.Log.Debug("suc: ", gsuc)
		s.Log.Debug("stat: ", stat)
	}
}

// AccessToken LinkedIn AccessToken
func (s *LinkedInSignin) AccessToken(redirectURI string, code string) *TokenResponse {
	var rtn TokenResponse

	return &rtn
}

// // Callback CallbackCallback
// func (s *LinkedInSignin) Callback(w http.ResponseWriter) *TokenParams {
// 	var rtn TokenParams

// 	//check scope in callback and if good, set:
// 	//LinkedInTokanParams.AuthCode = code from callback

// 	return &rtn
// }
