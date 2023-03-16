package delegates

import (
	"bytes"
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

// MCDelegate MCDelegate
type MCDelegate struct {
	proxy       px.Proxy
	Log         lg.Log
	RestURL     string
	APIKey      string
	APIAdminKey string
}

// New New
func (d *MCDelegate) New() Delegate {
	var p px.GoProxy
	d.proxy = &p
	return d
}

// ResponseID ResponseID
type ResponseID struct {
	ID      int64  `json:"id"`
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Response Response
type Response struct {
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// buildRequest buildRequest
func (d *MCDelegate) buildRequest(method string, url string, aJSON []byte, keyType string) (*http.Request, error) {
	var req *http.Request
	var err error
	var surl = d.RestURL + url
	if method == http.MethodPost || method == http.MethodPut {
		req, err = http.NewRequest(method, surl, bytes.NewBuffer(aJSON))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, surl, nil)
	}
	if keyType == adminKey {
		req.Header.Set("apiAdminKey", d.APIAdminKey)
	} else {
		req.Header.Set("apiKey", d.APIKey)
	}
	return req, err
}
