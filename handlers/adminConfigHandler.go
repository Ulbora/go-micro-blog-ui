package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

// ConfigPage ConfigPage
type ConfigPage struct {
	Title    string
	Desc     string
	KeyWords string
	Config   *mcd.Config
	SiteData *SiteData
}

// GetConfig GetConfig
func (h *MCHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetConfig")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetConfig", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetConfig: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {

			conf := h.Delegate.GetConfig()

			var cf ConfigPage
			cf.Title = h.Title
			cf.Desc = h.Desc
			cf.KeyWords = h.KeyWords
			cf.Config = conf

			h.AdminTemplates.ExecuteTemplate(w, adminConfigPage, &cf)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// UpdateConfig UpdateConfig
func (h *MCHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in UpdateConfig")
	cid := r.FormValue("cid")
	h.Log.Debug("cid in UpdateConfig: ", cid)
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in UpdateConfig: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			cnf := h.processConfig(r)
			res := h.Delegate.UpdateConfig(cnf)
			if !res.Success {
				h.Log.Debug("UpdateConfig success: ", res.Success)
				h.Log.Debug("UpdateConfig code: ", res.Code)
			}
			http.Redirect(w, r, adminConfigRt, http.StatusFound)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processConfig(r *http.Request) *mcd.Config {
	var rtn mcd.Config
	cidStr := r.FormValue("cid")
	h.Log.Debug("cid in processConfig: ", r.FormValue("cid"))
	h.Log.Debug("allowAutoPostOn: ", r.FormValue("allowAutoPost"))
	h.Log.Debug("allowAutoCommentOn: ", r.FormValue("allowAutoComment"))
	cid, _ := strconv.ParseInt(cidStr, 10, 64)
	allowAutoPostOn := r.FormValue("allowAutoPost")
	allowAutoCommentOn := r.FormValue("allowAutoComment")

	rtn.ID = cid
	if allowAutoPostOn == "on" {
		rtn.AllowAutoPost = true
	}
	if allowAutoCommentOn == "on" {
		rtn.AllowAutoComment = true
	}
	return &rtn
}
