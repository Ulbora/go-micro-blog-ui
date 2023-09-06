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
	//UserRole name
	UserRole = "user"
	//AdminRole name
	AdminRole = "admin"
)

// Delegate Delegate
type Delegate interface {
	AddUser(u *User) *ResponseID
	UpdateUser(u *User) *Response
	GetUser(email string) *User
	GetUserByID(id int64) *User
	GetUserList() *[]User
	GetUnActivatedUserList() *[]User
	GetBannedUserList() *[]User
	EnableUser(u *User) *Response
	DisableUser(u *User) *Response
	DisableUserForCause(u *User) *Response
	ReactivateUser(u *User) *Response

	AddRole(r *Role) *ResponseID
	GetRole(name string) *Role
	GetRoleList() *[]Role
	DeleteRole(rid int64) *Response

	AddBlog(b *Blog) *ResponseID
	UpdateBlog(b *Blog) *Response
	GetBlog(bid int64) *Blog
	GetBlogByName(name string, start int64, end int64) *[]Blog
	GetAdminBlogByName(name string, start int64, end int64) *[]Blog
	GetBlogList(start int64, end int64) *[]Blog
	GetAdminBlogList(start int64, end int64) *[]Blog
	ActivateBlog(b *Blog) *Response
	DeActivateBlog(b *Blog) *Response
	DeleteBlog(bid int64) *Response

	AddLike(l *Like) *ResponseID  // AddLike()
	RemoveLike(l *Like) *Response // RemoveLike()
	ViewLikes(bid int64) *[]Like  // ViewLikes()

	AddComment(c *Comment) *ResponseID                                   // AddComment()
	UpdateComment(c *Comment) *Response                                  // UpdateComment()
	GetCommentList(blogID int64, start int64, end int64) *[]Comment      // GetCommentList()
	GetAdminCommentList(blogID int64, start int64, end int64) *[]Comment // GetCommentAdminList()
	ActivateComment(c *Comment) *Response                                // ActivateComment()
	DeActivateComment(c *Comment) *Response                              // DectivateComment()

	AddUserAuth(a *UserAuth) *ResponseID                              // AddUserAuth()
	GetUserAuthList(userID int64, start int64, end int64) *[]UserAuth // GetUserAuthList()

	UpdateConfig(c *Config) *Response // UpdateConfig()
	GetConfig() *Config               // GetConfig()

	AddRule(r *Rule) *ResponseID
	UpdateRule(r *Rule) *Response
	GetRule() *Rule

	AddTerms(r *Terms) *ResponseID
	UpdateTerms(r *Terms) *Response
	GetTerms() *Terms

	AddAbout(r *About) *ResponseID
	UpdateAbout(r *About) *Response
	GetAbout() *About

	AddHome(r *Home) *ResponseID
	UpdateHome(r *Home) *Response
	GetHome() *Home

	AddPrivacyPolicy(r *PrivacyPolicy) *ResponseID
	UpdatePrivacyPolicy(r *PrivacyPolicy) *Response
	GetPrivacyPolicy() *PrivacyPolicy
}
