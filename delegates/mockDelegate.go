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

// MockDelegate MockDelegate
type MockDelegate struct {
	MockAddUserResp *ResponseID

	MockUpdateUserResp *Response

	MockRole *Role

	MockAddRoleResp *ResponseID

	MockRoleList *[]Role

	MockDeleteRoleResp *Response

	MockAddUserAuthResp *ResponseID

	MockUserAuthList *[]UserAuth

	MockUser *User

	MockBlogList *[]Blog
	MockBlog     *Blog

	MockCommentList *[]Comment

	MockLikeList *[]Like

	MockAddBlogRes *ResponseID

	MockUpdateBlogRes *Response

	MockAddCommentRes *ResponseID

	MockUpdateCommentRes *Response

	MockAddLikeRes *ResponseID

	MockRemoveLikeResp *Response

	MockActivateRes *Response

	MockDeActivateRes *Response

	MockDeleteBlogRes *Response

	MockActivateCommentRes *Response

	MockDeActivateCommentRes *Response

	MockUserList *[]User

	MockUnactivatedUserList *[]User

	MockBannedUserList *[]User

	MockEnableUser *Response

	MockReactivateBannedUser *Response

	MockDisableUser *Response

	MockDisableUserForCause *Response

	MockConfig *Config

	MockUpdateConfigRes *Response

	MockAddRuleResp *ResponseID

	MockUpdateRuleResp *Response

	MockRule *Rule

	MockAddTermsResp *ResponseID

	MockUpdateTermsResp *Response

	MockTerms *Terms

	MockAddAboutResp *ResponseID

	MockUpdateAboutResp *Response

	MockAbout *About
}

// New New
func (d *MockDelegate) New() Delegate {
	return d
}

// AddUser AddUser
func (d *MockDelegate) AddUser(u *User) *ResponseID {
	return d.MockAddUserResp
}

// UpdateUser UpdateUser
func (d *MockDelegate) UpdateUser(u *User) *Response {
	return d.MockUpdateUserResp

}

// GetUser GetUser
func (d *MockDelegate) GetUser(email string) *User {
	return d.MockUser

}

// GetUserByID GetUserByID
func (d *MockDelegate) GetUserByID(id int64) *User {
	return d.MockUser

}

// GetUserList GetUserList
func (d *MockDelegate) GetUserList() *[]User {
	return d.MockUserList

}

// GetUnActivatedUserList GetUnActivatedUserList
func (d *MockDelegate) GetUnActivatedUserList() *[]User {
	return d.MockUnactivatedUserList

}

// GetBannedUserList GetBannedUserList
func (d *MockDelegate) GetBannedUserList() *[]User {
	return d.MockBannedUserList

}

// EnableUser EnableUser
func (d *MockDelegate) EnableUser(u *User) *Response {
	return d.MockEnableUser
}

// DisableUser DisableUser
func (d *MockDelegate) DisableUser(u *User) *Response {
	return d.MockDisableUser

}

// DisableUserForCause DisableUserForCause
func (d *MockDelegate) DisableUserForCause(u *User) *Response {
	return d.MockDisableUserForCause

}

// ReactivateUser ReactivateUser
func (d *MockDelegate) ReactivateUser(u *User) *Response {
	return d.MockReactivateBannedUser

}

// AddRole AddRole
func (d *MockDelegate) AddRole(r *Role) *ResponseID {
	return d.MockAddRoleResp
}

// GetRole GetRole
func (d *MockDelegate) GetRole(name string) *Role {
	return d.MockRole
}

// GetRoleList GetRoleList
func (d *MockDelegate) GetRoleList() *[]Role {
	return d.MockRoleList
}

// DeleteRole DeleteRole
func (d *MockDelegate) DeleteRole(rid int64) *Response {
	return d.MockDeleteRoleResp
}

// AddBlog AddBlog
func (d *MockDelegate) AddBlog(b *Blog) *ResponseID {
	return d.MockAddBlogRes
}

// UpdateBlog UpdateBlog
func (d *MockDelegate) UpdateBlog(b *Blog) *Response {
	return d.MockUpdateBlogRes
}

// GetBlog GetBlog
func (d *MockDelegate) GetBlog(bid int64) *Blog {
	return d.MockBlog
}

// GetBlogByName GetBlogByName
func (d *MockDelegate) GetBlogByName(name string, start int64, end int64) *[]Blog {
	return d.MockBlogList
}

// GetAdminBlogByName GetAdminBlogByName
func (d *MockDelegate) GetAdminBlogByName(name string, start int64, end int64) *[]Blog {
	return d.MockBlogList
}

// GetBlogList GetBlogList
func (d *MockDelegate) GetBlogList(start int64, end int64) *[]Blog {
	return d.MockBlogList
}

// GetAdminBlogList GetAdminBlogList
func (d *MockDelegate) GetAdminBlogList(start int64, end int64) *[]Blog {
	return d.MockBlogList
}

// ActivateBlog ActivateBlog
func (d *MockDelegate) ActivateBlog(b *Blog) *Response {
	return d.MockActivateRes
}

// DeActivateBlog DeActivateBlog
func (d *MockDelegate) DeActivateBlog(b *Blog) *Response {
	return d.MockDeActivateRes
}

// DeleteBlog DeleteBlog
func (d *MockDelegate) DeleteBlog(bid int64) *Response {
	return d.MockDeleteBlogRes
}

// AddLike AddLike
func (d *MockDelegate) AddLike(l *Like) *ResponseID {
	return d.MockAddLikeRes
}

// RemoveLike RemoveLike
func (d *MockDelegate) RemoveLike(l *Like) *Response {
	return d.MockRemoveLikeResp
}

// ViewLikes ViewLikes
func (d *MockDelegate) ViewLikes(bid int64) *[]Like {
	return d.MockLikeList
}

// AddComment AddComment
func (d *MockDelegate) AddComment(c *Comment) *ResponseID {
	return d.MockAddCommentRes
}

// UpdateComment UpdateComment
func (d *MockDelegate) UpdateComment(c *Comment) *Response {
	return d.MockUpdateCommentRes
}

// GetCommentList GetCommentList
func (d *MockDelegate) GetCommentList(blogID int64, start int64, end int64) *[]Comment {
	return d.MockCommentList
}

// GetAdminCommentList GetAdminCommentList
func (d *MockDelegate) GetAdminCommentList(blogID int64, start int64, end int64) *[]Comment {
	return d.MockCommentList
}

// ActivateComment ActivateComment
func (d *MockDelegate) ActivateComment(c *Comment) *Response {
	return d.MockActivateCommentRes
}

// DeActivateComment DeActivateComment
func (d *MockDelegate) DeActivateComment(c *Comment) *Response {
	return d.MockDeActivateCommentRes
}

// AddUserAuth AddUserAuth
func (d *MockDelegate) AddUserAuth(a *UserAuth) *ResponseID {
	return d.MockAddUserAuthResp
}

// GetUserAuthList GetUserAuthList
func (d *MockDelegate) GetUserAuthList(userID int64, start int64, end int64) *[]UserAuth {
	return d.MockUserAuthList
}

// UpdateConfig UpdateConfig
func (d *MockDelegate) UpdateConfig(c *Config) *Response {
	return d.MockUpdateConfigRes
}

// GetConfig GetConfig
func (d *MockDelegate) GetConfig() *Config {
	return d.MockConfig
}

// AddRule AddRule
func (d *MockDelegate) AddRule(r *Rule) *ResponseID {
	return d.MockAddRuleResp
}

// UpdateRule UpdateRule
func (d *MockDelegate) UpdateRule(r *Rule) *Response {
	return d.MockUpdateRuleResp
}

// GetRule GetRule
func (d *MockDelegate) GetRule() *Rule {
	return d.MockRule
}

// AddTerms AddTerms
func (d *MockDelegate) AddTerms(r *Terms) *ResponseID {
	return d.MockAddTermsResp
}

// UpdateTerms UpdateTerms
func (d *MockDelegate) UpdateTerms(r *Terms) *Response {
	return d.MockUpdateTermsResp
}

// GetTerms GetTerms
func (d *MockDelegate) GetTerms() *Terms {
	return d.MockTerms
}

// AddAbout AddAbout
func (d *MockDelegate) AddAbout(r *About) *ResponseID {
	return d.MockAddAboutResp
}

// UpdateAbout UpdateAbout
func (d *MockDelegate) UpdateAbout(r *About) *Response {
	return d.MockUpdateAboutResp
}

// GetAbout GetAbout
func (d *MockDelegate) GetAbout() *About {
	return d.MockAbout
}
