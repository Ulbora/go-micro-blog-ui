package signins

import (
	"net/http"
	"net/url"

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
	linkedInDefaultAuthURL     = "https://www.linkedin.com/oauth/v2/authorization"
	linkedInDefaultTokenURL    = "https://www.linkedin.com/oauth/v2/accessToken"
	linkedInDefaultRedirectURL = "http://localhost:8080/auth/linkedin/callback"
	linkedInUserInfoURL        = "https://api.linkedin.com/v2/userinfo"
	scope                      = "profile%20email%20w_member_social%20r_liteprofile%20r_emailaddress%20openid"
	//State for linkedIn calls
	State = "ieridf7fsf6dfs6"
)

// UserInfoResponse UserInfoResponse
type UserInfoResponse struct {
	Sub           string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	Email         string `json:"email"`
	PictureURL    string `json:"picture"`
}

// LinkedInSignin LinkedInSignin
type LinkedInSignin struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	UserInfoURL  string
	RedirectURI  string
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
	if s.RedirectURI == "" {
		s.RedirectURI = linkedInDefaultRedirectURL
	}
	if s.UserInfoURL == "" {
		s.UserInfoURL = linkedInUserInfoURL
	}
	var p px.GoProxy
	s.proxy = &p
	return s
}

// SetProxy SetProxy
func (s *LinkedInSignin) SetProxy(p px.Proxy) {
	s.proxy = p
}

// Authorization LinkedIn Authorization
func (s *LinkedInSignin) Authorization(w http.ResponseWriter, r *http.Request) {
	//var rtn Response
	var aURL = s.AuthURL + "?response_type=code&client_id=" +
		s.ClientID + "&redirect_uri=" + s.RedirectURI + "&state=" + State +
		"&scope=" + scope
	s.Log.Debug("url: ", aURL)
	http.Redirect(w, r, aURL, http.StatusFound)
}

// AccessToken LinkedIn AccessToken
func (s *LinkedInSignin) AccessToken(code string) *TokenResponse {
	var rtn TokenResponse
	var tURL = s.TokenURL
	s.Log.Debug("url: ", tURL)
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", s.ClientID)
	data.Set("client_secret", s.ClientSecret)
	data.Set("redirect_uri", s.RedirectURI)
	brq, err := buildFormRequest(http.MethodPost, tURL, data)
	if err == nil {
		suc, stat := s.proxy.Do(brq, &rtn)
		s.Log.Debug("suc: ", suc)
		s.Log.Debug("stat: ", stat)
	}

	return &rtn
}

// GetUserInfo GetUserInfo
func (s *LinkedInSignin) GetUserInfo(tk string) any {
	var rtn UserInfoResponse
	var uiURL = s.UserInfoURL
	s.Log.Debug("uiURL: ", uiURL)

	urq, err := buildRequest(http.MethodGet, uiURL, nil)
	if err == nil {
		urq.Header.Set("Authorization", "Bearer "+tk)
		suc, stat := s.proxy.Do(urq, &rtn)
		s.Log.Debug("suc: ", suc)
		s.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetUserPicture GetUserPicture
func (s *LinkedInSignin) GetUserPicture(url string) []byte {
	var rtn []byte
	uprq, err := buildRequest(http.MethodGet, url, nil)
	if err == nil {
		suc, stat, b := s.proxy.DoNonJSON(uprq)
		rtn = b
		s.Log.Debug("suc: ", suc)
		s.Log.Debug("stat: ", stat)
	}
	return rtn
}
