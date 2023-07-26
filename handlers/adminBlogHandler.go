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

// GetAdminBlogList GetAdminBlogList
func (h *MCHandler) GetAdminBlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetAdminBlogList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			blogList := h.Delegate.GetAdminBlogList(0, maxPosts)
			h.Log.Debug("blogCnt: ", len(*blogList))
			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				bp.MyEmail = uemail.(string)
			}
			// var isAdmin = s.Get("isAdmin")
			h.Log.Debug("isAdmin: ", isAdmin)
			if isAdmin == true {
				bp.IsAdmin = true
			}
			var blst []Blog
			for i := range *blogList {
				var wg sync.WaitGroup
				var bb Blog
				bb.Blog = &(*blogList)[i]
				txt, err := b64.StdEncoding.DecodeString(bb.Blog.Content)
				if err == nil {
					bb.Blog.Content = string(txt)
					bb.Blog.Content = strings.Replace(bb.Blog.Content, stripOut, "", -1)
					bb.Blog.Content = strings.Replace(bb.Blog.Content, stripOut2, "", -1)
					bb.Blog.Content = strings.Replace(bb.Blog.Content, stripOut3, "", -1)
					bb.Blog.Content = strings.Replace(bb.Blog.Content, stripOut4, "", -1)

					bb.TextHTML = template.HTML(bb.Blog.Content)
					//bb.TextHTML = strings.Replace(bb.TextHTML, stripOut, "")
					// h.Log.Debug("TextHTML: ", bb.TextHTML)
				}
				//var ccnt int
				//var lcnt int
				wg.Add(1)
				go func(bbb *Blog) {
					// wg.Add(1)
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
					h.Log.Debug("commentCnt: ", len(*cl))
					// bb.CommentCnt = len(*cl)
					//ccnt = len(*cl)
					bbb.CommentCnt = len(*cl)
					h.Log.Debug("Comments Done: ")
				}(&bb)
				wg.Add(1)
				go func(bbb *Blog) {
					// wg.Add(1)
					defer wg.Done()
					ll := h.Delegate.ViewLikes(bbb.Blog.ID)
					for _, l := range *ll {
						if l.UserID == bbb.Blog.UserID {
							bbb.Liked = true
						}
					}
					h.Log.Debug("likeCnt: ", len(*ll))
					// bb.LikeCnt = len(*ll)
					bbb.LikeCnt = len(*ll)
					h.Log.Debug("Views Done: ")
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

				wg.Wait()
				// h.Log.Debug("after wait: for ", i)
				// h.Log.Debug("comments : ", ccnt)
				// h.Log.Debug("likes : ", lcnt)
				// bb.CommentCnt = ccnt
				// bb.LikeCnt = lcnt

				blst = append(blst, bb)
				// h.Log.Debug("blog: ", bb)
			}
			bp.BlogList = &blst
			h.Log.Debug("after all waits")
			h.Log.Debug("template: ", *h.AdminTemplates)
			//res := h.Service.GetContentList(false)
			// sort.Slice(*res, func(p, q int) bool {
			// 	return (*res)[p].Title < (*res)[q].Title
			// })
			// h.Log.Debug("content in admin index sorted: ", *res)
			h.AdminTemplates.ExecuteTemplate(w, adminBlogListPage, &bp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
	//<object data="data:application/pdf;base64,YOURBASE64DATA" type="application/pdf"></object>
}

// ActivateBlog ActivateBlog
func (h *MCHandler) ActivateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Activate Blog")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in Activate Blog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in Activate Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			vars := mux.Vars(r)
			bidStr := vars["bid"]
			h.Log.Debug("bid in Activate Blog: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)
			//uemail := s.Get("userEmail")
			//if uemail != nil {
			//email := uemail.(string)
			//h.Log.Debug("getting Blog: ")
			//abb := h.Delegate.GetBlog(bid)
			//h.Log.Debug("Blog: ")

			//var lk mcd.Like
			//lk.BlogID = bid
			//lk.UserID = u.ID
			var abb mcd.Blog
			abb.ID = bid
			//abb.Active = true
			res := h.Delegate.ActivateBlog(&abb)
			if !res.Success {
				h.Log.Debug("add Activate Blog suc: ", res.Success)
				h.Log.Debug("add Activate Blog code: ", res.Code)
				//h.Delegate.RemoveLike(&lk)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)
			//}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DectivateBlog DectivateBlog
func (h *MCHandler) DectivateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Dectivate Blog")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in Dectivate Blog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in Dectivate Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			vars := mux.Vars(r)
			bidStr := vars["bid"]
			h.Log.Debug("bid: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)
			//uemail := s.Get("userEmail")
			//if uemail != nil {
			//email := uemail.(string)
			//dbb := h.Delegate.GetBlog(bid)

			var dbb mcd.Blog
			dbb.ID = bid
			//var lk mcd.Like
			//lk.BlogID = bid
			//lk.UserID = u.ID
			//dbb.Active = false
			res := h.Delegate.DeActivateBlog(&dbb)
			if !res.Success {
				h.Log.Debug("add Dectivate Blog suc: ", res.Success)
				h.Log.Debug("add Dectivate Blog code: ", res.Code)
				//h.Delegate.RemoveLike(&lk)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)
			//}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
