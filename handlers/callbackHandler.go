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
	"time"

	del "github.com/Ulbora/go-micro-blog-ui/delegates"
	s "github.com/Ulbora/go-micro-blog-ui/signins"
)

// LinkedInCallback LinkedInCallback
func (h *MCHandler) LinkedInCallback(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	h.Log.Debug("in linkedIn callback")

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	si := h.Signins["linkedIn"]
	h.Log.Debug("state: ", state)
	h.Log.Debug("code: ", code)
	if s.State == state {
		tk := si.AccessToken(code)
		h.Log.Debug("token: ", tk.AccessToken)
		uiRes := si.GetUserInfo(tk.AccessToken)
		if uiRes.(*s.UserInfoResponse).EmailVerified && uiRes.(*s.UserInfoResponse).Email != "" {
			uemail := uiRes.(*s.UserInfoResponse).Email
			fname := uiRes.(*s.UserInfoResponse).FirstName
			lname := uiRes.(*s.UserInfoResponse).LastName
			picture := uiRes.(*s.UserInfoResponse).PictureURL

			h.Log.Debug("uemail: ", uemail)
			h.Log.Debug("fname: ", fname)
			h.Log.Debug("lname: ", lname)
			h.Log.Debug("picture: ", picture)
			h.Log.Debug("token: ", tk.AccessToken)

			var usuc bool
			var urole int64
			user := h.Delegate.GetUser(uemail)
			if !user.Active && user.Email == "" {
				h.Log.Debug("create new user: ", uemail)
				role := h.Delegate.GetRole(del.UserRole)
				var nusr del.User
				nusr.Active = true
				nusr.Email = uemail
				nusr.FirstName = fname
				nusr.LastName = lname
				nusr.RoleID = role.ID
				image := si.GetUserPicture(picture)
				nusr.Image = image
				adusr := h.Delegate.AddUser(&nusr)
				if adusr.Success {
					usuc = true
					var usauth del.UserAuth
					usauth.AuthType = "LinkedIn"
					usauth.Entered = time.Now()
					usauth.UserID = adusr.ID
					h.Delegate.AddUserAuth(&usauth)
				}
			} else if user.Active {
				urole = user.RoleID
				usuc = true
				var usauth del.UserAuth
				usauth.AuthType = "LinkedIn"
				usauth.Entered = time.Now()
				usauth.UserID = user.ID
				h.Delegate.AddUserAuth(&usauth)
			}

			if usuc {
				sec, suc := h.getSession(r)
				if suc {
					arole := h.Delegate.GetRole(del.AdminRole)
					if urole == arole.ID {
						sec.Set("isAdmin", true)
					}
					sec.Set("loggedIn", true)
					sec.Set("userEmail", uemail)
					serr := sec.Save(w)
					h.Log.Debug("serr", serr)

				}
				h.Log.Debug("session suc in linkedIn callback", suc)
			}

		}
		http.Redirect(w, r, indexRt, http.StatusFound)
	} else {
		http.Redirect(w, r, loginRt, http.StatusFound)
	}

}

// GoogleSigninCallback GoogleSigninCallback
func (h *MCHandler) GoogleSigninCallback(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	h.Log.Debug("in GoogleSigninCallback callback")

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	si := h.Signins["googleOAuth2"]
	h.Log.Debug("state: ", state)
	h.Log.Debug("code: ", code)
	if s.GoogleState == state {
		tk := si.AccessToken(code)
		h.Log.Debug("token: ", tk.AccessToken)
		guiRes := si.GetUserInfo(tk.AccessToken)
		if guiRes.(*s.GoogleUserInfoResponse).EmailVerified && guiRes.(*s.GoogleUserInfoResponse).Email != "" {
			uemail := guiRes.(*s.GoogleUserInfoResponse).Email
			fname := guiRes.(*s.GoogleUserInfoResponse).FirstName
			lname := guiRes.(*s.GoogleUserInfoResponse).LastName
			picture := guiRes.(*s.GoogleUserInfoResponse).PictureURL

			h.Log.Debug("guemail: ", uemail)
			h.Log.Debug("gfname: ", fname)
			h.Log.Debug("glname: ", lname)
			h.Log.Debug("gpicture: ", picture)
			h.Log.Debug("gtoken: ", tk.AccessToken)

			var gusuc bool
			var urole int64
			guser := h.Delegate.GetUser(uemail)
			if !guser.Active && guser.Email == "" {
				h.Log.Debug("create new user: ", uemail)
				role := h.Delegate.GetRole(del.UserRole)
				var gnusr del.User
				gnusr.Active = true
				gnusr.Email = uemail
				gnusr.FirstName = fname
				gnusr.LastName = lname
				gnusr.RoleID = role.ID
				image := si.GetUserPicture(picture)
				gnusr.Image = image
				adusr := h.Delegate.AddUser(&gnusr)
				if adusr.Success {
					gusuc = true
					var usauth del.UserAuth
					usauth.AuthType = "GoogleOAuth2"
					usauth.Entered = time.Now()
					usauth.UserID = adusr.ID
					h.Delegate.AddUserAuth(&usauth)
				}
			} else if guser.Active {
				urole = guser.RoleID
				gusuc = true
				var usauth del.UserAuth
				usauth.AuthType = "GoogleOAuth2"
				usauth.Entered = time.Now()
				usauth.UserID = guser.ID
				h.Delegate.AddUserAuth(&usauth)
			}

			if gusuc {
				gsec, suc := h.getSession(r)
				if suc {
					arole := h.Delegate.GetRole(del.AdminRole)
					if urole == arole.ID {
						gsec.Set("isAdmin", true)
					}
					gsec.Set("loggedIn", true)
					gsec.Set("userEmail", uemail)
					serr := gsec.Save(w)
					h.Log.Debug("serr", serr)

				}
				h.Log.Debug("session suc in google callback", suc)
			}

		}
		http.Redirect(w, r, indexRt, http.StatusFound)
	} else {
		http.Redirect(w, r, loginRt, http.StatusFound)
	}

}
