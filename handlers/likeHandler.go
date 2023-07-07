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
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

// AddLike AddLike
func (h *MCHandler) AddLike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlogList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {
			vars := mux.Vars(r)
			bidStr := vars["bid"]
			h.Log.Debug("bid: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)
			uemail := s.Get("userEmail")
			if uemail != nil {
				email := uemail.(string)
				u := h.Delegate.GetUser(email)

				var lk mcd.Like
				lk.BlogID = bid
				lk.UserID = u.ID
				res := h.Delegate.AddLike(&lk)
				if !res.Success {
					h.Log.Debug("add like suc: ", res.Success)
					h.Log.Debug("add like code: ", res.Code)
					h.Delegate.RemoveLike(&lk)
				}
				http.Redirect(w, r, indexRt, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
