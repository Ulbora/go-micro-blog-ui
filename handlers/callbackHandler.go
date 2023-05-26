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
)

// LinkedInCallback LinkedInCallback
func (h *MCHandler) LinkedInCallback(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	h.Log.Info("in linkedIn callback")

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	//var state = vars["state"]
	//var code = vars["code"]
	h.Log.Info("state: ", state)
	h.Log.Info("code: ", code)
	w.WriteHeader(http.StatusOK)

}
