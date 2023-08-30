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

// About About
type About struct {
	About    *mcd.About
	TextHTML template.HTML
	// User        *mcd.User
	// UserImage   string
	// CommentCnt  int
	// CommentList *[]Comment
	// LikeCnt     int
	// Liked       bool
}

// AboutPage AboutPage
type AboutPage struct {
	Title    string
	Desc     string
	KeyWords string
	About    *About
	IsAdmin  bool
	SiteData *SiteData
}

// SetAboutPage SetAboutPage
func (h *MCHandler) SetAboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetAboutPage")
	sr, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := sr.Get("loggedIn")
		var isAdmin = sr.Get("isAdmin")
		h.Log.Debug("loggedIn in SetAboutPage: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			h.Log.Debug("template: ", h.Templates)
			about := h.Delegate.GetAbout()

			var upg AboutPage
			upg.Title = h.Title
			upg.Desc = h.Desc
			upg.KeyWords = h.KeyWords

			var rl About
			rl.About = about

			h.Log.Debug("Content: ", about.Content)

			txt, err := b64.StdEncoding.DecodeString(about.Content)
			h.Log.Debug("text: ", string(txt))
			h.Log.Debug("err: ", err)
			if err == nil {
				about.Content = string(txt)
				about.Content = strings.Replace(about.Content, stripOut, "", -1)
				about.Content = strings.Replace(about.Content, stripOut2, "", -1)
				about.Content = strings.Replace(about.Content, stripOut3, "", -1)
				about.Content = strings.Replace(about.Content, stripOut4, "", -1)

				rl.TextHTML = template.HTML(about.Content)
				h.Log.Debug("TextHTML: ", rl.TextHTML)
			}
			upg.About = &rl
			h.AdminTemplates.ExecuteTemplate(w, adminAddAboutPage, &upg)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetAbout GetAbout
func (h *MCHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlog")
	grs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetAbout", suc)
	if suc {
		loggedInAuth := grs.Get("loggedIn")
		h.Log.Debug("loggedIn in GetAbout: ", loggedInAuth)
		if loggedInAuth == true {

			about := h.Delegate.GetAbout()

			var grp AboutPage
			grp.Title = h.Title
			grp.Desc = h.Desc
			grp.KeyWords = h.KeyWords

			var rrr About
			rrr.About = about
			txt, err := b64.StdEncoding.DecodeString(about.Content)
			if err == nil {
				about.Content = string(txt)
				about.Content = strings.Replace(about.Content, stripOut, "", -1)
				about.Content = strings.Replace(about.Content, stripOut2, "", -1)
				about.Content = strings.Replace(about.Content, stripOut3, "", -1)
				about.Content = strings.Replace(about.Content, stripOut4, "", -1)

				rrr.TextHTML = template.HTML(about.Content)
				h.Log.Debug("TextHTML: ", rrr.TextHTML)
			}

			grp.About = &rrr

			h.Templates.ExecuteTemplate(w, aboutPage, &grp)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// SetAbout SetAbout
func (h *MCHandler) SetAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetAbout")
	rid := r.FormValue("id")
	h.Log.Debug("id in SetAbout: ", rid)
	srs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := srs.Get("loggedIn")
		var isAdmin = srs.Get("isAdmin")
		h.Log.Debug("loggedIn in SetAbout: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			rll := h.processAbout(r)
			var rsucc bool
			var code int64
			if rll.ID == 0 {
				res := h.Delegate.AddAbout(rll)
				rsucc = res.Success
				code = res.Code
			} else {
				res := h.Delegate.UpdateAbout(rll)
				rsucc = res.Success
				code = res.Code
			}
			if !rsucc {
				h.Log.Debug("SetAbout success: ", rsucc)
				h.Log.Debug("SetAbout code: ", code)
			}
			http.Redirect(w, r, showAbout, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processAbout(r *http.Request) *mcd.About {
	var rtn mcd.About
	idStr := r.FormValue("id")
	content := r.FormValue("content")
	h.Log.Debug("id in processAbout: ", r.FormValue("id"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	return &rtn
}
