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
				}
				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
					h.Log.Debug("commentCnt: ", len(*cl))
					bbb.CommentCnt = len(*cl)
					h.Log.Debug("Comments Done: ")
				}(&bb)
				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					ll := h.Delegate.ViewLikes(bbb.Blog.ID)
					for _, l := range *ll {
						if l.UserID == bbb.Blog.UserID {
							bbb.Liked = true
						}
					}
					h.Log.Debug("likeCnt: ", len(*ll))
					bbb.LikeCnt = len(*ll)
					h.Log.Debug("Views Done: ")
				}(&bb)

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
					h.Log.Debug("get user: ")
					bbb.User = u1
					bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
					h.Log.Debug("User Done: ")
				}(&bb)

				wg.Wait()

				blst = append(blst, bb)
			}
			bp.BlogList = &blst
			h.Log.Debug("after all waits")
			h.Log.Debug("template: ", *h.AdminTemplates)

			h.AdminTemplates.ExecuteTemplate(w, adminBlogListPage, &bp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
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

			var abb mcd.Blog
			abb.ID = bid
			res := h.Delegate.ActivateBlog(&abb)
			if !res.Success {
				h.Log.Debug("add Activate Blog suc: ", res.Success)
				h.Log.Debug("add Activate Blog code: ", res.Code)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DeactivateBlog DeactivateBlog
func (h *MCHandler) DeactivateBlog(w http.ResponseWriter, r *http.Request) {
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

			var dbb mcd.Blog
			dbb.ID = bid

			res := h.Delegate.DeActivateBlog(&dbb)
			if !res.Success {
				h.Log.Debug("add Dectivate Blog suc: ", res.Success)
				h.Log.Debug("add Dectivate Blog code: ", res.Code)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// DeleteBlog DeleteBlog
func (h *MCHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Delete Blog")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in Delete Blog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		var isAdmin = s.Get("isAdmin")
		h.Log.Debug("loggedIn in Delete Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			vars := mux.Vars(r)
			bidStr := vars["bid"]
			h.Log.Debug("bid: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)

			res := h.Delegate.DeleteBlog(bid)
			if !res.Success {
				h.Log.Debug("add Delete Blog suc: ", res.Success)
				h.Log.Debug("add Delete Blog code: ", res.Code)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
