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
	googleOAuth2DefaultAuthURL     = "https://accounts.google.com/o/oauth2/v2/auth"
	googleOAuth2DefaultTokenURL    = "https://oauth2.googleapis.com/token"
	googleOAuth2DefaultRedirectURL = "https://localhost:8080/signin-google/callback"
	googleOAuth2UserInfoURL        = "https://www.googleapis.com/oauth2/v2/userinfo?alt=json"
	googleScope                    = "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile openid"
	//State for linkedIn calls
	googleState = "ieridf7fsf6dfs65775"
)

// GoogleUserInfoResponse GoogleUserInfoResponse
type GoogleUserInfoResponse struct {
	Sub           string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	Email         string `json:"email"`
	PictureURL    string `json:"picture"`
}

// GoogleSignin GoogleSignin
type GoogleSignin struct {
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
func (s *GoogleSignin) New() Signin {
	if s.AuthURL == "" {
		s.AuthURL = googleOAuth2DefaultAuthURL
	}
	if s.TokenURL == "" {
		s.TokenURL = googleOAuth2DefaultTokenURL
	}
	if s.RedirectURI == "" {
		s.RedirectURI = googleOAuth2DefaultRedirectURL
	}
	if s.UserInfoURL == "" {
		s.UserInfoURL = googleOAuth2UserInfoURL
	}
	var p px.GoProxy
	s.proxy = &p
	return s
}

// SetProxy SetProxy
func (s *GoogleSignin) SetProxy(p px.Proxy) {
	s.proxy = p
}

// Authorization LinkedIn Authorization
func (s *GoogleSignin) Authorization(w http.ResponseWriter, r *http.Request) {
	//var rtn Response
	var aURL = s.AuthURL + "?response_type=code&client_id=" +
		s.ClientID + "&redirect_uri=" + s.RedirectURI + "&state=" + State +
		"&scope=" + scope
	s.Log.Debug("url: ", aURL)
	http.Redirect(w, r, aURL, http.StatusFound)
}

// AccessToken LinkedIn AccessToken
func (s *GoogleSignin) AccessToken(code string) *TokenResponse {
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
func (s *GoogleSignin) GetUserInfo(tk string) any {
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
func (s *GoogleSignin) GetUserPicture(url string) []byte {
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
