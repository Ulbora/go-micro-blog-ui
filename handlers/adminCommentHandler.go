package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"

	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

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

// GetAdminCommentList GetAdminCommentList
func (h *MCHandler) GetAdminCommentList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetAdminCommentList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetAdminCommentList", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetAdminCommentList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			bvars := mux.Vars(r)
			bidstr := bvars["bid"]
			bid, _ := strconv.ParseInt(bidstr, 10, 64)
			h.Log.Debug("blog id in edit", bid)
			var wg sync.WaitGroup
			bg := h.Delegate.GetBlog(bid)

			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				bp.MyEmail = uemail.(string)
			}
			var bb Blog
			bb.Blog = bg
			txt, err := b64.StdEncoding.DecodeString(bg.Content)
			if err == nil {
				bg.Content = string(txt)
				bg.Content = strings.Replace(bg.Content, stripOut, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut2, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut3, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut4, "", -1)

				bb.TextHTML = template.HTML(bg.Content)
				//bb.TextHTML = strings.Replace(bb.TextHTML, stripOut, "")
				h.Log.Debug("TextHTML: ", bb.TextHTML)
			}

			bp.Blog = &bb
			wg.Add(1)
			go func(bbb *Blog) {
				// wg.Add(1)
				defer wg.Done()
				cl := h.Delegate.GetAdminCommentList(bbb.Blog.ID, 0, maxComments)
				h.Log.Debug("commentCnt: ", len(*cl))
				// bb.CommentCnt = len(*cl)
				//ccnt = len(*cl)
				var cmtLst []Comment
				for i := range *cl {
					cmt := (*cl)[i]
					var c Comment
					c.Comment = &cmt
					u1 := h.Delegate.GetUserByID(cmt.UserID)
					c.User = u1
					c.UserImage = b64.StdEncoding.EncodeToString(u1.Image)
					cmtLst = append(cmtLst, c)
				}
				bbb.CommentList = &cmtLst
				bbb.CommentCnt = len(*cl)
				h.Log.Debug("Comments Done: ")
			}(&bb)

			wg.Add(1)
			go func(bbb *Blog) {
				// wg.Add(1)
				defer wg.Done()
				u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
				h.Log.Debug("get user: ")
				// bb.CommentCnt = len(*cl)
				//ccnt = len(*cl)
				bbb.User = u1
				bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
				h.Log.Debug("User Done: ")
			}(&bb)

			// wg.Add(1)
			// go func(bbb *Blog) {
			// 	// wg.Add(1)
			// 	defer wg.Done()
			// 	u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
			// 	h.Log.Debug("get user: ")
			// 	// bb.CommentCnt = len(*cl)
			// 	//ccnt = len(*cl)
			// 	bbb.User = u1
			// 	bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
			// 	h.Log.Debug("User Done: ")
			// }(&bb)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminBlogPage, &bp)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// ActivateComment ActivateComment
func (h *MCHandler) ActivateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Activate Comment")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in Activate Comment", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in Activate Comment: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			vars := mux.Vars(r)
			cidStr := vars["cid"]
			bidStr := vars["bid"]
			h.Log.Debug("cid in Activate Comment: ", cidStr)
			cid, _ := strconv.ParseInt(cidStr, 10, 64)
			//bid, _ := strconv.ParseInt(bidStr, 10, 64)
			//uemail := s.Get("userEmail")
			//if uemail != nil {
			//email := uemail.(string)
			//h.Log.Debug("getting Blog: ")
			//abb := h.Delegate.GetBlog(bid)
			//h.Log.Debug("Blog: ")

			//var lk mcd.Like
			//lk.BlogID = bid
			//lk.UserID = u.ID
			var acc mcd.Comment
			acc.ID = cid
			//abb.Active = true
			res := h.Delegate.ActivateComment(&acc)
			if !res.Success {
				h.Log.Debug("add Activate Comment suc: ", res.Success)
				h.Log.Debug("add Activate Comment code: ", res.Code)
				//h.Delegate.RemoveLike(&lk)
			}
			http.Redirect(w, r, viesAdminPostRt+"/"+bidStr, http.StatusFound)
			//}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DectivateComment DectivateComment
func (h *MCHandler) DectivateComment(w http.ResponseWriter, r *http.Request) {

}
