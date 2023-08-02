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
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"

	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

// UserPage UserPage
type UserPage struct {
	Title     string
	Desc      string
	KeyWords  string
	User      *mcd.User
	Users     *[]User
	UserImage string
}

// User User
type User struct {
	User      *mcd.User
	UserImage string
}

// GetUserByIDPage GetUserByIDPage
func (h *MCHandler) GetUserByIDPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetUserByIDPage")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetUserByIDPage", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			bvars := mux.Vars(r)
			uidstr := bvars["uid"]
			uid, _ := strconv.ParseInt(uidstr, 10, 64)
			h.Log.Debug("user id", uid)

			user := h.Delegate.GetUserByID(uid)

			var up UserPage
			up.Title = h.Title
			up.Desc = h.Desc
			up.KeyWords = h.KeyWords
			up.User = user
			up.UserImage = b64.StdEncoding.EncodeToString(up.User.Image)

			h.AdminTemplates.ExecuteTemplate(w, adminUserPage, &up)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// GetUnactivatedUserList GetUnactivatedUserList
func (h *MCHandler) GetUnactivatedUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetUnactivatedUserList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetUnactivatedUserList", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetUnactivatedUserList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			// bvars := mux.Vars(r)
			// ststr := bvars["start"]
			// str, _ := strconv.ParseInt(ststr, 10, 64)
			// h.Log.Debug("page start", str)

			// edstr := bvars["end"]
			// end, _ := strconv.ParseInt(edstr, 10, 64)
			// h.Log.Debug("page end", end)

			users := h.Delegate.GetUnActivatedUserList()
			var ul []User
			for i := range *users {
				var u User
				u.User = &(*users)[i]
				u.UserImage = b64.StdEncoding.EncodeToString(u.User.Image)
				ul = append(ul, u)
			}

			var up UserPage
			up.Title = h.Title
			up.Desc = h.Desc
			up.KeyWords = h.KeyWords
			up.Users = &ul

			h.AdminTemplates.ExecuteTemplate(w, adminUnactivatedUserPage, &up)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// GetBannedUserList GetBannedUserList
func (h *MCHandler) GetBannedUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBannedUserList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBannedUserList", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBannedUserList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			// bvars := mux.Vars(r)
			// ststr := bvars["start"]
			// str, _ := strconv.ParseInt(ststr, 10, 64)
			// h.Log.Debug("page start", str)

			// edstr := bvars["end"]
			// end, _ := strconv.ParseInt(edstr, 10, 64)
			// h.Log.Debug("page end", end)

			users := h.Delegate.GetBannedUserList()
			var ul []User
			for i := range *users {
				var u User
				u.User = &(*users)[i]
				u.UserImage = b64.StdEncoding.EncodeToString(u.User.Image)
				ul = append(ul, u)
			}

			var up UserPage
			up.Title = h.Title
			up.Desc = h.Desc
			up.KeyWords = h.KeyWords
			up.Users = &ul

			h.AdminTemplates.ExecuteTemplate(w, adminBannedUserPage, &up)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
