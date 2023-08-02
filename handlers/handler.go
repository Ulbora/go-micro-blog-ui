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

import "net/http"

const (
	linkedInSystem = "linkedIn"
	googleOAuth2   = "googleOAuth2"

	//pages
	loginPage    = "login.html"
	blogListPage = "blogList.html"
	blogPage     = "blog.html"
	addBlogPage  = "addBlog.html"
	editBlogPage = "editBlog.html"

	//Admin Pages
	adminBlogListPage        = "adminBlogList.html"
	adminBlogPage            = "adminBlog.html"
	adminUserPage            = "adminUser.html"
	adminUnactivatedUserPage = "adminUnactivatedUser.html"
	adminBannedUserPage      = "adminBannedUser.html"

	//routes
	indexRt  = "/"
	loginRt  = "/login"
	viewPost = "/viewPost/"

	//admin routes
	adminBlogListRt = "/adminPostList"
	viesAdminPostRt = "/adminViewPost"

	maxPosts    = 200
	maxComments = 1000
)

// Handler Handler
type Handler interface {
	//---------------for admin user ---------------------------------

	GetUserByIDPage(w http.ResponseWriter, r *http.Request)
	// GetUserList(w http.ResponseWriter, r *http.Request)
	GetUnactivatedUserList(w http.ResponseWriter, r *http.Request)
	GetBannedUserList(w http.ResponseWriter, r *http.Request)
	// EnableUser(w http.ResponseWriter, r *http.Request)
	// DisableUser(w http.ResponseWriter, r *http.Request)
	// DisableUserForCause(w http.ResponseWriter, r *http.Request)
	// ReinstateBannedUser(w http.ResponseWriter, r *http.Request)

	// AddRolePage(w http.ResponseWriter, r *http.Request)
	// AddRole(w http.ResponseWriter, r *http.Request)
	// GetRole(w http.ResponseWriter, r *http.Request)
	// GetRoleList(w http.ResponseWriter, r *http.Request)
	// DeleteRole(w http.ResponseWriter, r *http.Request)

	// GetConfig(w http.ResponseWriter, r *http.Request)
	// UpdateConfig(w http.ResponseWriter, r *http.Request)

	// GetUserAuthList(w http.ResponseWriter, r *http.Request)

	GetAdminBlogList(w http.ResponseWriter, r *http.Request)
	ActivateBlog(w http.ResponseWriter, r *http.Request)
	DeactivateBlog(w http.ResponseWriter, r *http.Request)

	GetAdminCommentList(w http.ResponseWriter, r *http.Request)
	ActivateComment(w http.ResponseWriter, r *http.Request)
	DeactivateComment(w http.ResponseWriter, r *http.Request)

	//----------------------------for users-------------------------
	//user
	LoginUserPage(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)

	// //callbacks
	LinkedInCallback(w http.ResponseWriter, r *http.Request)
	GoogleSigninCallback(w http.ResponseWriter, r *http.Request)

	// //===== CreateNewUserPage(w http.ResponseWriter, r *http.Request)
	// //===== CreateNewUser(w http.ResponseWriter, r *http.Request)

	// //===== ActivateNewUser(w http.ResponseWriter, r *http.Request)

	// UpdateUserPage(w http.ResponseWriter, r *http.Request)
	// UpdateUser(w http.ResponseWriter, r *http.Request)
	// //=====GetUser(w http.ResponseWriter, r *http.Request)

	// //user blog
	AddBlogPage(w http.ResponseWriter, r *http.Request)
	AddBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlogPage(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
	// GetBlogByName(w http.ResponseWriter, r *http.Request)
	GetBlogList(w http.ResponseWriter, r *http.Request)

	AddLike(w http.ResponseWriter, r *http.Request)
	// RemoveLike(w http.ResponseWriter, r *http.Request)
	// ViewLikes(w http.ResponseWriter, r *http.Request)

	// AddCommentPage(w http.ResponseWriter, r *http.Request)
	AddComment(w http.ResponseWriter, r *http.Request)
	// //maybe getcomment for update?
	// GetComment(w http.ResponseWriter, r *http.Request)
	// UpdateComment(w http.ResponseWriter, r *http.Request)
	// GetCommentList(w http.ResponseWriter, r *http.Request)
}
