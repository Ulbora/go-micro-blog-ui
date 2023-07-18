package delegates

// MockDelegate MockDelegate
type MockDelegate struct {
	AddUserID      int64
	AddUserSuccess bool

	GetUserActive bool
	GetUserEmail  string

	GetRoleID   int64
	GetRoleName string

	AddUserAuthID      int64
	AddUserAuthSuccess bool

	MockUser *User

	MockBlogList *[]Blog
	MockBlog     *Blog

	MockCommentList *[]Comment

	MockLikeList *[]Like

	MockAddBlogRes *ResponseID

	MockAddCommentRes *ResponseID

	MockAddLikeRes *ResponseID
}

// New New
func (d *MockDelegate) New() Delegate {
	return d
}

// AddUser AddUser
func (d *MockDelegate) AddUser(u *User) *ResponseID {
	var rtn ResponseID
	rtn.ID = d.AddUserID
	rtn.Success = d.AddUserSuccess
	return &rtn
}

// UpdateUser UpdateUser
func (d *MockDelegate) UpdateUser(u *User) *Response {
	return nil

}

// GetUser GetUser
func (d *MockDelegate) GetUser(email string) *User {
	var rtn User
	rtn.Active = d.GetUserActive
	rtn.Email = d.GetUserEmail
	return &rtn

}

// GetUserByID GetUserByID
func (d *MockDelegate) GetUserByID(id int64) *User {
	return d.MockUser

}

// GetUserList GetUserList
func (d *MockDelegate) GetUserList() *[]User {
	return nil

}

// GetUnActivatedUserList GetUnActivatedUserList
func (d *MockDelegate) GetUnActivatedUserList() *[]User {
	return nil

}

// GetBannedUserList GetBannedUserList
func (d *MockDelegate) GetBannedUserList() *[]User {
	return nil

}

// EnableUser EnableUser
func (d *MockDelegate) EnableUser(u *User) *Response {
	return nil

}

// DisableUser DisableUser
func (d *MockDelegate) DisableUser(u *User) *Response {
	return nil

}

// DisableUserForCause DisableUserForCause
func (d *MockDelegate) DisableUserForCause(u *User) *Response {
	return nil

}

// ReactivateUser ReactivateUser
func (d *MockDelegate) ReactivateUser(u *User) *Response {
	return nil

}

// AddRole AddRole
func (d *MockDelegate) AddRole(r *Role) *ResponseID {
	return nil
}

// GetRole GetRole
func (d *MockDelegate) GetRole(name string) *Role {
	var rtn Role
	rtn.ID = d.GetRoleID
	rtn.Name = d.GetRoleName
	return &rtn
}

// GetRoleList GetRoleList
func (d *MockDelegate) GetRoleList() *[]Role {
	return nil
}

// DeleteRole DeleteRole
func (d *MockDelegate) DeleteRole(rid int64) *Response {
	return nil
}

// AddBlog AddBlog
func (d *MockDelegate) AddBlog(b *Blog) *ResponseID {
	return d.MockAddBlogRes
}

// UpdateBlog UpdateBlog
func (d *MockDelegate) UpdateBlog(b *Blog) *Response {
	return nil
}

// GetBlog GetBlog
func (d *MockDelegate) GetBlog(bid int64) *Blog {
	return d.MockBlog
}

// GetBlogByName GetBlogByName
func (d *MockDelegate) GetBlogByName(name string, start int64, end int64) *[]Blog {
	return nil
}

// GetBlogList GetBlogList
func (d *MockDelegate) GetBlogList(start int64, end int64) *[]Blog {
	return d.MockBlogList
}

// GetAdminBlogList GetAdminBlogList
func (d *MockDelegate) GetAdminBlogList(start int64, end int64) *[]Blog {
	return nil
}

// ActivateBlog ActivateBlog
func (d *MockDelegate) ActivateBlog(b *Blog) *Response {
	return nil
}

// DeActivateBlog DeActivateBlog
func (d *MockDelegate) DeActivateBlog(b *Blog) *Response {
	return nil
}

// AddLike AddLike
func (d *MockDelegate) AddLike(l *Like) *ResponseID {
	return d.MockAddLikeRes
}

// RemoveLike RemoveLike
func (d *MockDelegate) RemoveLike(l *Like) *Response {
	return nil
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
	return nil
}

// GetCommentList GetCommentList
func (d *MockDelegate) GetCommentList(blogID int64, start int64, end int64) *[]Comment {
	return d.MockCommentList
}

// GetAdminCommentList GetAdminCommentList
func (d *MockDelegate) GetAdminCommentList(blogID int64, start int64, end int64) *[]Comment {
	return nil
}

// ActivateComment ActivateComment
func (d *MockDelegate) ActivateComment(c *Comment) *Response {
	return nil
}

// DeActivateComment DeActivateComment
func (d *MockDelegate) DeActivateComment(c *Comment) *Response {
	return nil
}

// AddUserAuth AddUserAuth
func (d *MockDelegate) AddUserAuth(a *UserAuth) *ResponseID {
	var rtn ResponseID
	rtn.ID = d.AddUserAuthID
	rtn.Success = d.AddUserAuthSuccess

	return &rtn
}

// GetUserAuthList GetUserAuthList
func (d *MockDelegate) GetUserAuthList(userID int64, start int64, end int64) *[]UserAuth {
	return nil
}

// UpdateConfig UpdateConfig
func (d *MockDelegate) UpdateConfig(c *Config) *Response {
	return nil
}

// GetConfig GetConfig
func (d *MockDelegate) GetConfig() *Config {
	return nil
}
