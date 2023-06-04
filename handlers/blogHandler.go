package handlers

import (
	"fmt"
	"net/http"
)

// GetBlogList GetBlogList
func (h *MCHandler) GetBlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlogList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {
			//h.Log.Debug("template: ", h.AdminTemplates)
			//res := h.Service.GetContentList(false)
			// sort.Slice(*res, func(p, q int) bool {
			// 	return (*res)[p].Title < (*res)[q].Title
			// })
			// h.Log.Debug("content in admin index sorted: ", *res)
			h.Templates.ExecuteTemplate(w, blogListPage, nil)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
