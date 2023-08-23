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
	gs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetUserByIDPage", suc)
	if suc {
		loggedInAuth := gs.Get("loggedIn")
		var isAdmin = gs.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			guvars := mux.Vars(r)
			uidstr := guvars["uid"]
			uid, _ := strconv.ParseInt(uidstr, 10, 64)
			h.Log.Debug("user id", uid)

			user := h.Delegate.GetUserByID(uid)

			var gup UserPage
			gup.Title = h.Title
			gup.Desc = h.Desc
			gup.KeyWords = h.KeyWords
			gup.User = user
			gup.UserImage = b64.StdEncoding.EncodeToString(gup.User.Image)

			h.AdminTemplates.ExecuteTemplate(w, adminUserPage, &gup)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// GetUnactivatedUserList GetUnactivatedUserList
func (h *MCHandler) GetUnactivatedUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetUnactivatedUserList")
	uas, suc := h.getSession(r)
	h.Log.Debug("session suc in GetUnactivatedUserList", suc)
	if suc {
		loggedInAuth := uas.Get("loggedIn")
		var isAdmin = uas.Get("isAdmin")
		h.Log.Debug("loggedIn in GetUnactivatedUserList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {

			users := h.Delegate.GetUnActivatedUserList()
			var ual []User
			for i := range *users {
				var u User
				u.User = &(*users)[i]
				u.UserImage = b64.StdEncoding.EncodeToString(u.User.Image)
				ual = append(ual, u)
			}

			var uap UserPage
			uap.Title = h.Title
			uap.Desc = h.Desc
			uap.KeyWords = h.KeyWords
			uap.Users = &ual

			h.AdminTemplates.ExecuteTemplate(w, adminUnactivatedUserPage, &uap)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// GetBannedUserList GetBannedUserList
func (h *MCHandler) GetBannedUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBannedUserList")
	gbs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBannedUserList", suc)
	if suc {
		loggedInAuth := gbs.Get("loggedIn")
		var isAdmin = gbs.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBannedUserList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {

			busers := h.Delegate.GetBannedUserList()
			var ul []User
			for i := range *busers {
				var u User
				u.User = &(*busers)[i]
				u.UserImage = b64.StdEncoding.EncodeToString(u.User.Image)
				ul = append(ul, u)
			}

			var bup UserPage
			bup.Title = h.Title
			bup.Desc = h.Desc
			bup.KeyWords = h.KeyWords
			bup.Users = &ul

			h.AdminTemplates.ExecuteTemplate(w, adminBannedUserPage, &bup)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// EnableUser EnableUser
func (h *MCHandler) EnableUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in EnableUser")
	es, suc := h.getSession(r)
	h.Log.Debug("session suc in EnableUser", suc)
	if suc {
		loggedInAuth := es.Get("loggedIn")
		var isAdmin = es.Get("isAdmin")
		h.Log.Debug("loggedIn in EnableUser: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			evars := mux.Vars(r)
			uidStr := evars["uid"]
			h.Log.Debug("uid in EnableUser: ", uidStr)
			uid, _ := strconv.ParseInt(uidStr, 10, 64)

			var eubb mcd.User
			eubb.ID = uid
			eres := h.Delegate.EnableUser(&eubb)
			if !eres.Success {
				h.Log.Debug("add EnableUser suc: ", eres.Success)
				h.Log.Debug("add EnableUser code: ", eres.Code)
			}
			http.Redirect(w, r, adminUserRt+"/"+uidStr, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DisableUser DisableUser
func (h *MCHandler) DisableUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in DisableUser")
	das, suc := h.getSession(r)
	h.Log.Debug("session suc in DisableUser", suc)
	if suc {
		loggedInAuth := das.Get("loggedIn")
		var isAdmin = das.Get("isAdmin")
		h.Log.Debug("loggedIn in DisableUser: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			davars := mux.Vars(r)
			uidStr := davars["uid"]
			h.Log.Debug("uid in DisableUser: ", uidStr)
			uid, _ := strconv.ParseInt(uidStr, 10, 64)

			var dubb mcd.User
			dubb.ID = uid
			dares := h.Delegate.DisableUser(&dubb)
			if !dares.Success {
				h.Log.Debug("add DisableUser suc: ", dares.Success)
				h.Log.Debug("add DisableUser code: ", dares.Code)
			}
			http.Redirect(w, r, adminUserRt+"/"+uidStr, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DisableUserForCause DisableUserForCause
func (h *MCHandler) DisableUserForCause(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in DisableUserForCause")
	dds, suc := h.getSession(r)
	h.Log.Debug("session suc in DisableUserForCause", suc)
	if suc {
		loggedInAuth := dds.Get("loggedIn")
		var isAdmin = dds.Get("isAdmin")
		h.Log.Debug("loggedIn in DisableUserForCause: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			dddvars := mux.Vars(r)
			uidStr := dddvars["uid"]
			h.Log.Debug("uid in DisableUserForCause: ", uidStr)
			uid, _ := strconv.ParseInt(uidStr, 10, 64)

			var dddubb mcd.User
			dddubb.ID = uid
			dddres := h.Delegate.DisableUserForCause(&dddubb)
			if !dddres.Success {
				h.Log.Debug("add DisableUserForCause suc: ", dddres.Success)
				h.Log.Debug("add DisableUserForCause code: ", dddres.Code)
			}
			http.Redirect(w, r, adminUserRt+"/"+uidStr, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// ReinstateBannedUser ReinstateBannedUser
func (h *MCHandler) ReinstateBannedUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in ReinstateBannedUser")
	ris, suc := h.getSession(r)
	h.Log.Debug("session suc in ReinstateBannedUser", suc)
	if suc {
		loggedInAuth := ris.Get("loggedIn")
		var isAdmin = ris.Get("isAdmin")
		h.Log.Debug("loggedIn in ReinstateBannedUser: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			rivars := mux.Vars(r)
			uidStr := rivars["uid"]
			h.Log.Debug("uid in ReinstateBannedUser: ", uidStr)
			uid, _ := strconv.ParseInt(uidStr, 10, 64)

			var riubb mcd.User
			riubb.ID = uid
			rires := h.Delegate.ReactivateUser(&riubb)
			if !rires.Success {
				h.Log.Debug("add ReinstateBannedUser suc: ", rires.Success)
				h.Log.Debug("add ReinstateBannedUser code: ", rires.Code)
			}
			http.Redirect(w, r, adminUserRt+"/"+uidStr, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
