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
	"net/http"

	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

// LinkedInCallback LinkedInCallback
func (h *MCHandler) LinkedInCallback(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	h.Log.Info("in linkedIn callback")

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	si := h.Signins["linkedIn"]
	h.Log.Info("state: ", state)
	h.Log.Info("code: ", code)
	if s.State == state {
		tk := si.AccessToken(code)
		h.Log.Info("token: ", tk.AccessToken)
		sec, suc := h.getSession(r)
		if suc {
			sec.Set("linkedInToken", &tk)
			serr := sec.Save(w)
			h.Log.Debug("serr", serr)
		}
		h.Log.Debug("session suc", suc)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
