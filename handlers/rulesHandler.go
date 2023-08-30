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

// Rule Rule
type Rule struct {
	Rule     *mcd.Rule
	TextHTML template.HTML
	// User        *mcd.User
	// UserImage   string
	// CommentCnt  int
	// CommentList *[]Comment
	// LikeCnt     int
	// Liked       bool
}

// RulePage RulePage
type RulePage struct {
	Title    string
	Desc     string
	KeyWords string
	Rule     *Rule
	IsAdmin  bool
	SiteData *SiteData
}

// SetRulesPage SetRulesPage
func (h *MCHandler) SetRulesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetRulesPage")
	sr, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := sr.Get("loggedIn")
		var isAdmin = sr.Get("isAdmin")
		h.Log.Debug("loggedIn in AddBlog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			h.Log.Debug("template: ", h.Templates)
			rule := h.Delegate.GetRule()

			var upg RulePage
			upg.Title = h.Title
			upg.Desc = h.Desc
			upg.KeyWords = h.KeyWords

			var rl Rule
			rl.Rule = rule

			h.Log.Debug("Content: ", rule.Content)
			// upg.Rule = &rl
			// uemail := sr.Get("userEmail")
			// if uemail != nil {
			// 	pd.UserEmail = uemail.(string)
			// }
			txt, err := b64.StdEncoding.DecodeString(rule.Content)
			h.Log.Debug("text: ", string(txt))
			h.Log.Debug("err: ", err)
			if err == nil {
				rule.Content = string(txt)
				rule.Content = strings.Replace(rule.Content, stripOut, "", -1)
				rule.Content = strings.Replace(rule.Content, stripOut2, "", -1)
				rule.Content = strings.Replace(rule.Content, stripOut3, "", -1)
				rule.Content = strings.Replace(rule.Content, stripOut4, "", -1)

				rl.TextHTML = template.HTML(rule.Content)
				h.Log.Debug("TextHTML: ", rl.TextHTML)
			}
			upg.Rule = &rl
			h.AdminTemplates.ExecuteTemplate(w, adminAddRulePage, &upg)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetRules GetRules
func (h *MCHandler) GetRules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlog")
	grs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := grs.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {

			rg := h.Delegate.GetRule()

			var grp RulePage
			grp.Title = h.Title
			grp.Desc = h.Desc
			grp.KeyWords = h.KeyWords

			var rrr Rule
			rrr.Rule = rg
			txt, err := b64.StdEncoding.DecodeString(rg.Content)
			if err == nil {
				rg.Content = string(txt)
				rg.Content = strings.Replace(rg.Content, stripOut, "", -1)
				rg.Content = strings.Replace(rg.Content, stripOut2, "", -1)
				rg.Content = strings.Replace(rg.Content, stripOut3, "", -1)
				rg.Content = strings.Replace(rg.Content, stripOut4, "", -1)

				rrr.TextHTML = template.HTML(rg.Content)
				h.Log.Debug("TextHTML: ", rrr.TextHTML)
			}

			grp.Rule = &rrr

			h.Templates.ExecuteTemplate(w, rulesPage, &grp)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// SetRules SetRules
func (h *MCHandler) SetRules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetRules")
	rid := r.FormValue("id")
	h.Log.Debug("id in SetRules: ", rid)
	srs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := srs.Get("loggedIn")
		var isAdmin = srs.Get("isAdmin")
		h.Log.Debug("loggedIn in SetRules: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			rll := h.processRule(r)
			var rsucc bool
			var code int64
			if rll.ID == 0 {
				res := h.Delegate.AddRule(rll)
				rsucc = res.Success
				code = res.Code
			} else {
				res := h.Delegate.UpdateRule(rll)
				rsucc = res.Success
				code = res.Code
			}
			if !rsucc {
				h.Log.Debug("SetRules success: ", rsucc)
				h.Log.Debug("SetRules code: ", code)
			}
			http.Redirect(w, r, showRules, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processRule(r *http.Request) *mcd.Rule {
	var rtn mcd.Rule
	idStr := r.FormValue("id")
	content := r.FormValue("content")
	h.Log.Debug("id in processRule: ", r.FormValue("id"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	return &rtn
}
