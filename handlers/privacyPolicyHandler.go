package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
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

// PrivacyPolicy PrivacyPolicy
type PrivacyPolicy struct {
	PrivacyPolicy *mcd.PrivacyPolicy
	TextHTML      template.HTML
	// User        *mcd.User
	// UserImage   string
	// CommentCnt  int
	// CommentList *[]Comment
	// LikeCnt     int
	// Liked       bool
}

// PrivacyPolicyPage PrivacyPolicyPage
type PrivacyPolicyPage struct {
	Title         string
	Desc          string
	KeyWords      string
	PrivacyPolicy *PrivacyPolicy
	IsAdmin       bool
	SiteData      *SiteData
}

// SetPrivacyPolicyPage SetPrivacyPolicyPage
func (h *MCHandler) SetPrivacyPolicyPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetPrivacyPolicyPage")
	sr, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := sr.Get("loggedIn")
		var isAdmin = sr.Get("isAdmin")
		h.Log.Debug("loggedIn in SetPrivacyPolicyPage: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			h.Log.Debug("template: ", h.Templates)
			pp := h.Delegate.GetPrivacyPolicy()

			var upg PrivacyPolicyPage
			upg.Title = h.Title
			upg.Desc = h.Desc
			upg.KeyWords = h.KeyWords

			var rl PrivacyPolicy
			rl.PrivacyPolicy = pp

			h.Log.Debug("Content: ", pp.Content)

			txt, err := b64.StdEncoding.DecodeString(pp.Content)
			h.Log.Debug("text: ", string(txt))
			h.Log.Debug("err: ", err)
			if err == nil {
				pp.Content = string(txt)
				pp.Content = strings.Replace(pp.Content, stripOut, "", -1)
				pp.Content = strings.Replace(pp.Content, stripOut2, "", -1)
				pp.Content = strings.Replace(pp.Content, stripOut3, "", -1)
				pp.Content = strings.Replace(pp.Content, stripOut4, "", -1)

				rl.TextHTML = template.HTML(pp.Content)
				h.Log.Debug("TextHTML: ", rl.TextHTML)
			}
			upg.PrivacyPolicy = &rl
			h.AdminTemplates.ExecuteTemplate(w, adminAddPrivacyPolicyPage, &upg)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetPrivacyPolicy GetPrivacyPolicy
func (h *MCHandler) GetPrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetPrivacyPolicy")
	grs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetPrivacyPolicy", suc)
	if suc {
		loggedInAuth := grs.Get("loggedIn")
		h.Log.Debug("loggedIn in GetPrivacyPolicy: ", loggedInAuth)
		// if loggedInAuth == true {

		pp := h.Delegate.GetPrivacyPolicy()

		var grp PrivacyPolicyPage
		grp.Title = h.Title
		grp.Desc = h.Desc
		grp.KeyWords = h.KeyWords

		var rrr PrivacyPolicy
		rrr.PrivacyPolicy = pp
		txt, err := b64.StdEncoding.DecodeString(pp.Content)
		if err == nil {
			pp.Content = string(txt)
			pp.Content = strings.Replace(pp.Content, stripOut, "", -1)
			pp.Content = strings.Replace(pp.Content, stripOut2, "", -1)
			pp.Content = strings.Replace(pp.Content, stripOut3, "", -1)
			pp.Content = strings.Replace(pp.Content, stripOut4, "", -1)

			rrr.TextHTML = template.HTML(pp.Content)
			h.Log.Debug("TextHTML: ", rrr.TextHTML)
		}

		grp.PrivacyPolicy = &rrr

		h.Templates.ExecuteTemplate(w, privacyPolicyPage, &grp)

		// }
		// else {
		// 	http.Redirect(w, r, loginRt, http.StatusFound)
		// }
	}
}

// SetPrivacyPolicy SetPrivacyPolicy
func (h *MCHandler) SetPrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetPrivacyPolicy")
	rid := r.FormValue("id")
	h.Log.Debug("id in SetPrivacyPolicy: ", rid)
	srs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := srs.Get("loggedIn")
		var isAdmin = srs.Get("isAdmin")
		h.Log.Debug("loggedIn in SetPrivacyPolicy: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			rll := h.processPrivacyPolicy(r)
			var rsucc bool
			var code int64
			if rll.ID == 0 {
				res := h.Delegate.AddPrivacyPolicy(rll)
				rsucc = res.Success
				code = res.Code
			} else {
				res := h.Delegate.UpdatePrivacyPolicy(rll)
				rsucc = res.Success
				code = res.Code
			}
			if !rsucc {
				h.Log.Debug("SetPrivacyPolicy success: ", rsucc)
				h.Log.Debug("SetPrivacyPolicy code: ", code)
			}
			http.Redirect(w, r, showPrivacyPolicy, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processPrivacyPolicy(r *http.Request) *mcd.PrivacyPolicy {
	var rtn mcd.PrivacyPolicy
	idStr := r.FormValue("id")
	content := r.FormValue("content")
	h.Log.Debug("id in processPrivacyPolicy: ", r.FormValue("id"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	return &rtn
}
