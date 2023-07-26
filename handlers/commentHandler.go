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

// AddComment AddComment
func (h *MCHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in AddComment")
	userEmail := r.FormValue("userEmail")
	bidStr := r.FormValue("bid")
	comment := r.FormValue("comment")

	h.Log.Debug("userEmail in AddComment: ", userEmail)
	h.Log.Debug("bidStr in AddComment: ", bidStr)
	h.Log.Debug("comment in AddComment: ", comment)
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in addBlog: ", loggedInAuth)
		if loggedInAuth == true {
			email, cmt := h.processComment(r)
			h.Log.Debug("email in addComment: ", email)
			h.Log.Debug("comment in addComment: ", cmt.Text)
			h.Log.Debug("bid in addComment: ", cmt.BlogID)
			u := h.Delegate.GetUser(email)
			cmt.UserID = u.ID
			res := h.Delegate.AddComment(cmt)
			h.Log.Debug("addComment success: ", res.Success)
			h.Log.Debug("addComment code: ", res.Code)
			http.Redirect(w, r, viewPost+bidStr, http.StatusFound)

		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processComment(r *http.Request) (string, *mcd.Comment) {
	var rtn mcd.Comment
	bidStr := r.FormValue("bid")
	bid, _ := strconv.ParseInt(bidStr, 10, 64)
	email := r.FormValue("userEmail")
	comment := r.FormValue("comment")
	h.Log.Debug("name in processComment: ", r.FormValue("comment"))

	//rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	rtn.BlogID = bid
	rtn.Text = comment

	return email, &rtn
}
