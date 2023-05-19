package handlers

import "net/http"

// Handler Handler
type Handler interface {
	//---------------for admin user ---------------------------------

	GetUserList(w http.ResponseWriter, r *http.Request)
	GetUnactivatedUserList(w http.ResponseWriter, r *http.Request)
	GetBannedUserList(w http.ResponseWriter, r *http.Request)
	EnableUser(w http.ResponseWriter, r *http.Request)
	DisableUser(w http.ResponseWriter, r *http.Request)
	DisableUserForCause(w http.ResponseWriter, r *http.Request)
	ReinstateBannedUser(w http.ResponseWriter, r *http.Request)

	AddRolePage(w http.ResponseWriter, r *http.Request)
	AddRole(w http.ResponseWriter, r *http.Request)
	GetRole(w http.ResponseWriter, r *http.Request)
	GetRoleList(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)

	GetConfig(w http.ResponseWriter, r *http.Request)
	UpdateConfig(w http.ResponseWriter, r *http.Request)

	GetUserAuthList(w http.ResponseWriter, r *http.Request)

	GetAdminBlogList(w http.ResponseWriter, r *http.Request)
	ActivateBlog(w http.ResponseWriter, r *http.Request)
	DectivateBlog(w http.ResponseWriter, r *http.Request)

	GetCommentAdminList(w http.ResponseWriter, r *http.Request)
	ActivateComment(w http.ResponseWriter, r *http.Request)
	DectivateComment(w http.ResponseWriter, r *http.Request)

	//----------------------------for users-------------------------
	//user
	LoginUserPage(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)

	CreateNewUserPage(w http.ResponseWriter, r *http.Request)
	CreateNewUser(w http.ResponseWriter, r *http.Request)

	ActivateNewUser(w http.ResponseWriter, r *http.Request)

	UpdateUserPage(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	//GetUser(w http.ResponseWriter, r *http.Request)

	//user blog
	AddBlogPage(w http.ResponseWriter, r *http.Request)
	AddBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlogPage(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
	GetBlogByName(w http.ResponseWriter, r *http.Request)
	GetBlogList(w http.ResponseWriter, r *http.Request)

	AddLike(w http.ResponseWriter, r *http.Request)
	RemoveLike(w http.ResponseWriter, r *http.Request)
	ViewLikes(w http.ResponseWriter, r *http.Request)

	AddComment(w http.ResponseWriter, r *http.Request)
	//maybe getcomment for update?
	UpdateComment(w http.ResponseWriter, r *http.Request)
	GetCommentList(w http.ResponseWriter, r *http.Request)
}
