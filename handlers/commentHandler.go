package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

// AddComment AddComment
func (h *MCHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in AddComment")
	userEmail := r.FormValue("userEmail")
	bidStr := r.FormValue("bid")
	comment := r.FormValue("comment")

	h.Log.Debug("userEmail in AddComment: ", userEmail)
	h.Log.Debug("bidStr in AddComment: ", bidStr)
	h.Log.Debug("comment in AddComment: ", comment)
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in addBlog: ", loggedInAuth)
		if loggedInAuth == true {
			email, cmt := h.processComment(r)
			h.Log.Debug("email in addComment: ", email)
			h.Log.Debug("comment in addComment: ", cmt.Text)
			h.Log.Debug("bid in addComment: ", cmt.BlogID)
			u := h.Delegate.GetUser(email)
			cmt.UserID = u.ID
			res := h.Delegate.AddComment(cmt)
			h.Log.Debug("addComment success: ", res.Success)
			h.Log.Debug("addComment code: ", res.Code)
			http.Redirect(w, r, viewPost+bidStr, http.StatusFound)

		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processComment(r *http.Request) (string, *mcd.Comment) {
	var rtn mcd.Comment
	bidStr := r.FormValue("bid")
	bid, _ := strconv.ParseInt(bidStr, 10, 64)
	email := r.FormValue("userEmail")
	comment := r.FormValue("comment")
	h.Log.Debug("name in processComment: ", r.FormValue("comment"))

	//rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	rtn.BlogID = bid
	rtn.Text = comment

	return email, &rtn
}
