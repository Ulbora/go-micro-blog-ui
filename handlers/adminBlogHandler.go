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
	abs, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := abs.Get("loggedIn")
		var isAdmin = abs.Get("isAdmin")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			ablogList := h.Delegate.GetAdminBlogList(0, maxPosts)
			h.Log.Debug("blogCnt: ", len(*ablogList))
			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			uemail := abs.Get("userEmail")
			if uemail != nil {
				bp.MyEmail = uemail.(string)
			}
			// var isAdmin = s.Get("isAdmin")
			h.Log.Debug("isAdmin: ", isAdmin)
			if isAdmin == true {
				bp.IsAdmin = true
			}
			var blst []Blog
			for i := range *ablogList {
				var wg sync.WaitGroup
				var abb Blog
				abb.Blog = &(*ablogList)[i]
				txt, err := b64.StdEncoding.DecodeString(abb.Blog.Content)
				if err == nil {
					abb.Blog.Content = string(txt)
					abb.Blog.Content = strings.Replace(abb.Blog.Content, stripOut, "", -1)
					abb.Blog.Content = strings.Replace(abb.Blog.Content, stripOut2, "", -1)
					abb.Blog.Content = strings.Replace(abb.Blog.Content, stripOut3, "", -1)
					abb.Blog.Content = strings.Replace(abb.Blog.Content, stripOut4, "", -1)

					abb.TextHTML = template.HTML(abb.Blog.Content)
				}
				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
					h.Log.Debug("commentCnt: ", len(*cl))
					bbb.CommentCnt = len(*cl)
					h.Log.Debug("Comments Done: ")
				}(&abb)
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
				}(&abb)

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
					h.Log.Debug("get user: ")
					bbb.User = u1
					bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
					h.Log.Debug("User Done: ")
				}(&abb)

				wg.Wait()

				blst = append(blst, abb)
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

// SearchAdminBlogList SearchAdminBlogList
func (h *MCHandler) SearchAdminBlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SearchAdminBlogList")
	aas, suc := h.getSession(r)
	h.Log.Debug("session suc in SearchAdminBlogList", suc)
	if suc {
		loggedInAuth := aas.Get("loggedIn")
		var isAdmin = aas.Get("isAdmin")
		h.Log.Debug("loggedIn in SearchAdminBlogList: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			name := r.FormValue("name")
			sblogList := h.Delegate.GetAdminBlogByName(name, 0, maxPosts)
			h.Log.Debug("blogCnt: ", len(*sblogList))
			var sbp BlogPage
			sbp.Title = h.Title
			sbp.Desc = h.Desc
			sbp.KeyWords = h.KeyWords
			uemail := aas.Get("userEmail")
			if uemail != nil {
				sbp.MyEmail = uemail.(string)
			}
			var isAdmin = aas.Get("isAdmin")
			h.Log.Debug("isAdmin: ", isAdmin)
			if isAdmin == true {
				sbp.IsAdmin = true
			}
			var ablst []Blog
			for i := range *sblogList {
				var wg sync.WaitGroup
				var asbb Blog
				asbb.Blog = &(*sblogList)[i]
				txt, err := b64.StdEncoding.DecodeString(asbb.Blog.Content)
				if err == nil {
					asbb.Blog.Content = string(txt)
					asbb.Blog.Content = strings.Replace(asbb.Blog.Content, stripOut, "", -1)
					asbb.Blog.Content = strings.Replace(asbb.Blog.Content, stripOut2, "", -1)
					asbb.Blog.Content = strings.Replace(asbb.Blog.Content, stripOut3, "", -1)
					asbb.Blog.Content = strings.Replace(asbb.Blog.Content, stripOut4, "", -1)

					asbb.TextHTML = template.HTML(asbb.Blog.Content)
					h.Log.Debug("TextHTML: ", asbb.TextHTML)
				}

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
					h.Log.Debug("commentCnt: ", len(*cl))
					bbb.CommentCnt = len(*cl)
					h.Log.Debug("Comments Done: ")
				}(&asbb)
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
				}(&asbb)

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
					h.Log.Debug("get user: ")
					bbb.User = u1
					bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
					h.Log.Debug("User Done: ")
				}(&asbb)

				wg.Wait()

				ablst = append(ablst, asbb)
			}
			sbp.BlogList = &ablst
			h.Log.Debug("after all waits")

			h.AdminTemplates.ExecuteTemplate(w, adminBlogListPage, &sbp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// ActivateBlog ActivateBlog
func (h *MCHandler) ActivateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Activate Blog")
	aas, suc := h.getSession(r)
	h.Log.Debug("session suc in Activate Blog", suc)
	if suc {
		loggedInAuth := aas.Get("loggedIn")
		var isAdmin = aas.Get("isAdmin")
		h.Log.Debug("loggedIn in Activate Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			aavars := mux.Vars(r)
			bidStr := aavars["bid"]
			h.Log.Debug("bid in Activate Blog: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)

			var aabb mcd.Blog
			aabb.ID = bid
			res := h.Delegate.ActivateBlog(&aabb)
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
	ads, suc := h.getSession(r)
	h.Log.Debug("session suc in Dectivate Blog", suc)
	if suc {
		loggedInAuth := ads.Get("loggedIn")
		var isAdmin = ads.Get("isAdmin")
		h.Log.Debug("loggedIn in Dectivate Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			advars := mux.Vars(r)
			bidStr := advars["bid"]
			h.Log.Debug("bid: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)

			var adbb mcd.Blog
			adbb.ID = bid

			res := h.Delegate.DeActivateBlog(&adbb)
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
	adds, suc := h.getSession(r)
	h.Log.Debug("session suc in Delete Blog", suc)
	if suc {
		loggedInAuth := adds.Get("loggedIn")
		var isAdmin = adds.Get("isAdmin")
		h.Log.Debug("loggedIn in Delete Blog: ", loggedInAuth)
		if loggedInAuth == true && isAdmin == true {
			addvars := mux.Vars(r)
			bidStr := addvars["bid"]
			h.Log.Debug("bid: ", bidStr)
			bid, _ := strconv.ParseInt(bidStr, 10, 64)

			addres := h.Delegate.DeleteBlog(bid)
			if !addres.Success {
				h.Log.Debug("add Delete Blog suc: ", addres.Success)
				h.Log.Debug("add Delete Blog code: ", addres.Code)
			}
			http.Redirect(w, r, adminBlogListRt, http.StatusFound)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}
