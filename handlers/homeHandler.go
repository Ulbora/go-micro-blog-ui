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

// Home Home
type Home struct {
	Home     *mcd.Home
	TextHTML template.HTML
	// User        *mcd.User
	// UserImage   string
	// CommentCnt  int
	// CommentList *[]Comment
	// LikeCnt     int
	// Liked       bool
}

// HomePage HomePage
type HomePage struct {
	Title    string
	Desc     string
	KeyWords string
	Home     *Home
	IsAdmin  bool
	SiteData *SiteData
}

// SetHomePage SetHomePage
func (h *MCHandler) SetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetHomePage")
	sr, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := sr.Get("loggedIn")
		var isAdmin = sr.Get("isAdmin")
		h.Log.Debug("loggedIn in SetHomePage: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			h.Log.Debug("template: ", h.Templates)
			home := h.Delegate.GetHome()

			var upg HomePage
			upg.Title = h.Title
			upg.Desc = h.Desc
			upg.KeyWords = h.KeyWords

			var rl Home
			rl.Home = home

			h.Log.Debug("Content: ", home.Content)

			txt, err := b64.StdEncoding.DecodeString(home.Content)
			h.Log.Debug("text: ", string(txt))
			h.Log.Debug("err: ", err)
			if err == nil {
				home.Content = string(txt)
				home.Content = strings.Replace(home.Content, stripOut, "", -1)
				home.Content = strings.Replace(home.Content, stripOut2, "", -1)
				home.Content = strings.Replace(home.Content, stripOut3, "", -1)
				home.Content = strings.Replace(home.Content, stripOut4, "", -1)

				rl.TextHTML = template.HTML(home.Content)
				h.Log.Debug("TextHTML: ", rl.TextHTML)
			}
			upg.Home = &rl
			h.AdminTemplates.ExecuteTemplate(w, adminAddHomePage, &upg)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetHome GetHome
func (h *MCHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetHome")
	grs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetHome", suc)
	if suc {
		loggedInAuth := grs.Get("loggedIn")
		h.Log.Debug("loggedIn in GetHome: ", loggedInAuth)
		// if loggedInAuth == true {

		home := h.Delegate.GetHome()

		var grp HomePage
		grp.Title = h.Title
		grp.Desc = h.Desc
		grp.KeyWords = h.KeyWords

		var rrr Home
		rrr.Home = home
		txt, err := b64.StdEncoding.DecodeString(home.Content)
		if err == nil {
			home.Content = string(txt)
			home.Content = strings.Replace(home.Content, stripOut, "", -1)
			home.Content = strings.Replace(home.Content, stripOut2, "", -1)
			home.Content = strings.Replace(home.Content, stripOut3, "", -1)
			home.Content = strings.Replace(home.Content, stripOut4, "", -1)

			rrr.TextHTML = template.HTML(home.Content)
			h.Log.Debug("TextHTML: ", rrr.TextHTML)
		}

		grp.Home = &rrr

		h.Templates.ExecuteTemplate(w, homePage, &grp)

		// }
		// else {
		// 	http.Redirect(w, r, loginRt, http.StatusFound)
		// }
	}
}

// SetHome SetHome
func (h *MCHandler) SetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SetHome")
	rid := r.FormValue("id")
	h.Log.Debug("id in SetHome: ", rid)
	srs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := srs.Get("loggedIn")
		var isAdmin = srs.Get("isAdmin")
		h.Log.Debug("loggedIn in SetHome: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			rll := h.processHome(r)
			var rsucc bool
			var code int64
			if rll.ID == 0 {
				res := h.Delegate.AddHome(rll)
				rsucc = res.Success
				code = res.Code
			} else {
				res := h.Delegate.UpdateHome(rll)
				rsucc = res.Success
				code = res.Code
			}
			if !rsucc {
				h.Log.Debug("SetHome success: ", rsucc)
				h.Log.Debug("SetHome code: ", code)
			}
			http.Redirect(w, r, showHome, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processHome(r *http.Request) *mcd.Home {
	var rtn mcd.Home
	idStr := r.FormValue("id")
	content := r.FormValue("content")
	h.Log.Debug("id in processHome: ", r.FormValue("id"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	return &rtn
}
