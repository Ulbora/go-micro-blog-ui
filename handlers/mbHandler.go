package handlers

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

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	lg "github.com/GolangToolKits/go-level-logger"
	gss "github.com/GolangToolKits/go-secure-sessions"
	m "github.com/Ulbora/go-micro-blog-ui/managers"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

// MCHandler MCHandler
type MCHandler struct {
	Log            lg.Log
	Manager        m.Manager
	APIKey         string
	APIAdminKey    string
	Signins        map[string]s.Signin
	SessionManager gss.SessionManager
	AdminTemplates *template.Template
	Templates      *template.Template
}

// New New
func (h *MCHandler) New() Handler {
	return h
}

func (h *MCHandler) processAPIKey(r *http.Request) bool {
	var rtn bool
	apiKey := r.Header.Get("apiKey")
	h.Log.Debug("apiKey: ", apiKey)
	h.Log.Debug("h.APIKey: ", h.APIKey)
	if apiKey == h.APIKey {
		rtn = true
	}
	return rtn
}

func (h *MCHandler) processAPIAdminKey(r *http.Request) bool {
	var rtn bool
	apiAdminKey := r.Header.Get("apiAdminKey")
	h.Log.Debug("apiAdminKey: ", apiAdminKey)
	h.Log.Debug("h.APIAdminKey: ", h.APIAdminKey)
	if apiAdminKey == h.APIAdminKey {
		rtn = true
	}
	return rtn
}

// CheckContent CheckContent
func (h *MCHandler) checkContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

// SetContentType SetContentType
func (h *MCHandler) setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// ProcessBody ProcessBody
func (h *MCHandler) processBody(r *http.Request, obj any) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	h.Log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}

func (h *MCHandler) getSession(r *http.Request) (gss.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn gss.Session
	gob.Register(&s.TokenResponse{})

	if r != nil {
		sec := h.SessionManager.NewSession(r, "go-micro-blog-ui")
		if sec != nil {
			suc = true
			loggedInAuth := sec.Get("loggedIn")
			h.Log.Debug("loggedIn: ", loggedInAuth)
		}
		srtn = sec
	}
	//fmt.Println("exit getSession--------------------------------------------------")
	return srtn, suc
}
