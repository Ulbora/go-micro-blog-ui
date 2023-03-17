package delegates

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

const (
	adminKey = "admin"
	apiKey   = "api"
)

// Delegate Delegate
type Delegate interface {
	AddUser(u *User) *ResponseID
	UpdateUser(u *User) *Response
	GetUser(email string) *User
	GetUserList() *[]User
	EnableUser(u *User) *Response
	DisableUser(u *User) *Response

	AddRole(r *Role) *ResponseID
	GetRole(name string) *Role
	GetRoleList() *[]Role
	DeleteRole(rid int64) *Response

	AddBlog(b *Blog) *ResponseID
	UpdateBlog(b *Blog) *Response
	GetBlog(bid int64) *Blog
	GetBlogByName(name string, start int64, end int64) *[]Blog
	GetBlogList(start int64, end int64) *[]Blog
	GetAdminBlogList(start int64, end int64) *[]Blog
	ActivateBlog(b *Blog) *Response
	DeActivateBlog(b *Blog) *Response

	// AddLike()
	// RemoveLike()
	// ViewLikes()

	// AddComment()
	// UpdateComment()
	// GetCommentList()
	// GetCommentAdminList()
	// ActivateComment()
	// DectivateComment()

	// AddUserAuth()
	// GetUserAuthList()

	// UpdateConfig()
	// GetConfig()

}
