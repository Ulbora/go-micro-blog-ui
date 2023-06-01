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
	"fmt"
	"net/http"
)

// LoginUserPage LoginUserPage
func (h *MCHandler) LoginUserPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in loginUserPage")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in login: ", loggedInAuth)
		if loggedInAuth == nil || loggedInAuth == false {
			h.Log.Debug("template: ", h.AdminTemplates)
			//res := h.Service.GetContentList(false)
			// sort.Slice(*res, func(p, q int) bool {
			// 	return (*res)[p].Title < (*res)[q].Title
			// })
			// h.Log.Debug("content in admin index sorted: ", *res)
			h.Templates.ExecuteTemplate(w, loginPage, nil)
		} else {
			http.Redirect(w, r, indexPage, http.StatusFound)
		}
	}
}

// LoginUser LoginUser
func (h *MCHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	signingSystem := r.FormValue("signingSystem")
	if signingSystem == linkedInSystem {

	}
}
