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

// Terms Terms
type Terms struct {
	Terms    *mcd.Terms
	TextHTML template.HTML
	// User        *mcd.User
	// UserImage   string
	// CommentCnt  int
	// CommentList *[]Comment
	// LikeCnt     int
	// Liked       bool
}

// TermsPage TermsPage
type TermsPage struct {
	Title    string
	Desc     string
	KeyWords string
	Terms    *Terms
	IsAdmin  bool
	SiteData *SiteData
}

// SetTermsPage SetTermsPage
func (h *MCHandler) SetTermsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetTermsPage")
	sr, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := sr.Get("loggedIn")
		var isAdmin = sr.Get("isAdmin")
		h.Log.Debug("loggedIn in SetTermsPage: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			h.Log.Debug("template: ", h.Templates)
			terms := h.Delegate.GetTerms()

			var upg TermsPage
			upg.Title = h.Title
			upg.Desc = h.Desc
			upg.KeyWords = h.KeyWords

			var rl Terms
			rl.Terms = terms

			h.Log.Debug("Content: ", terms.Content)

			txt, err := b64.StdEncoding.DecodeString(terms.Content)
			h.Log.Debug("text: ", string(txt))
			h.Log.Debug("err: ", err)
			if err == nil {
				terms.Content = string(txt)
				terms.Content = strings.Replace(terms.Content, stripOut, "", -1)
				terms.Content = strings.Replace(terms.Content, stripOut2, "", -1)
				terms.Content = strings.Replace(terms.Content, stripOut3, "", -1)
				terms.Content = strings.Replace(terms.Content, stripOut4, "", -1)

				rl.TextHTML = template.HTML(terms.Content)
				h.Log.Debug("TextHTML: ", rl.TextHTML)
			}
			upg.Terms = &rl
			h.AdminTemplates.ExecuteTemplate(w, adminAddTermsPage, &upg)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetTerms GetTerms
func (h *MCHandler) GetTerms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetTerms")
	grs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetTerms", suc)
	if suc {
		loggedInAuth := grs.Get("loggedIn")
		h.Log.Debug("loggedIn in GetTerms: ", loggedInAuth)
		// if loggedInAuth == true {

		tg := h.Delegate.GetTerms()

		var grp TermsPage
		grp.Title = h.Title
		grp.Desc = h.Desc
		grp.KeyWords = h.KeyWords

		var rrr Terms
		rrr.Terms = tg
		txt, err := b64.StdEncoding.DecodeString(tg.Content)
		if err == nil {
			tg.Content = string(txt)
			tg.Content = strings.Replace(tg.Content, stripOut, "", -1)
			tg.Content = strings.Replace(tg.Content, stripOut2, "", -1)
			tg.Content = strings.Replace(tg.Content, stripOut3, "", -1)
			tg.Content = strings.Replace(tg.Content, stripOut4, "", -1)

			rrr.TextHTML = template.HTML(tg.Content)
			h.Log.Debug("TextHTML: ", rrr.TextHTML)
		}

		grp.Terms = &rrr

		h.Templates.ExecuteTemplate(w, termsPage, &grp)

		// }
		// else {
		// http.Redirect(w, r, loginRt, http.StatusFound)
		// }
	}
}

// SetTerms SetTerms
func (h *MCHandler) SetTerms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetTerms")
	rid := r.FormValue("id")
	h.Log.Debug("id in SetTerms: ", rid)
	srs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := srs.Get("loggedIn")
		var isAdmin = srs.Get("isAdmin")
		h.Log.Debug("loggedIn in SetTerms: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			tll := h.processTerms(r)
			var rsucc bool
			var code int64
			if tll.ID == 0 {
				res := h.Delegate.AddTerms(tll)
				rsucc = res.Success
				code = res.Code
			} else {
				res := h.Delegate.UpdateTerms(tll)
				rsucc = res.Success
				code = res.Code
			}
			if !rsucc {
				h.Log.Debug("SetTerms success: ", rsucc)
				h.Log.Debug("SetTerms code: ", code)
			}
			http.Redirect(w, r, showTerms, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processTerms(r *http.Request) *mcd.Terms {
	var rtn mcd.Terms
	idStr := r.FormValue("id")
	content := r.FormValue("content")
	h.Log.Debug("id in processTerms: ", r.FormValue("id"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	return &rtn
}
