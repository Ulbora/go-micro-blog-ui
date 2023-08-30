package handlers

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

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	b64 "encoding/base64"

	mux "github.com/GolangToolKits/grrt"
	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

const (
	stripOut  = "<input type=\"text\" data-formula=\"e=mc^2\" data-link=\"https://quilljs.com\" data-video=\"Embed URL\" placeholder=\"Embed URL\">"
	stripOut2 = "<input type=\"text\" data-formula=\"e=mc^2\" data-link=\"https://quilljs.com\" data-video=\"Embed URL\">"
	stripOut3 = "contenteditable=\"true\""
	stripOut4 = "<input type=\"text\" data-formula=\"e=mc^2\" data-link=\"https://quilljs.com\" data-video=\"Embed URL\" placeholder=\"https://quilljs.com\">"
)

// Blog Blog
type Blog struct {
	Blog        *mcd.Blog
	TextHTML    template.HTML
	User        *mcd.User
	UserImage   string
	CommentCnt  int
	CommentList *[]Comment
	LikeCnt     int
	Liked       bool
}

// Comment Comment
type Comment struct {
	Comment   *mcd.Comment
	User      *mcd.User
	UserImage string
}

// BlogPage BlogPage
type BlogPage struct {
	Title    string
	Desc     string
	KeyWords string
	BlogList *[]Blog
	Blog     *Blog
	MyEmail  string
	IsAdmin  bool
	SiteData *SiteData
}

// AddBlogPageData AddBlogPageData
type AddBlogPageData struct {
	Title     string
	Desc      string
	KeyWords  string
	UserEmail string
	SiteData  *SiteData
}

// GetBlogList GetBlogList
func (h *MCHandler) GetBlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlogList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {
			blogList := h.Delegate.GetBlogList(0, maxPosts)
			h.Log.Debug("blogCnt: ", len(*blogList))
			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				bp.MyEmail = uemail.(string)
			}
			var isAdmin = s.Get("isAdmin")
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
					h.Log.Debug("TextHTML: ", bb.TextHTML)
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

			h.Templates.ExecuteTemplate(w, blogListPage, &bp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// AddBlogPage AddBlogPage
func (h *MCHandler) AddBlogPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in AddBlogPage")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in AddBlog: ", loggedInAuth)
		if loggedInAuth == true {
			h.Log.Debug("template: ", h.Templates)
			var pd AddBlogPageData
			pd.Title = h.Title
			pd.Desc = h.Desc
			pd.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				pd.UserEmail = uemail.(string)
			}

			h.Templates.ExecuteTemplate(w, addBlogPage, &pd)
		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// AddBlog AddBlog
func (h *MCHandler) AddBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in AddBlog")
	name := r.FormValue("name")
	h.Log.Debug("name in addBlog: ", name)
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in addBlog: ", loggedInAuth)
		if loggedInAuth == true {
			email, blog := h.processBlog(r)
			h.Log.Debug("email in addBlog: ", email)
			h.Log.Debug("name in addBlog name: ", blog.Name)
			h.Log.Debug("blog in addBlog content: ", blog.Content)
			u := h.Delegate.GetUser(email)
			blog.UserID = u.ID
			res := h.Delegate.AddBlog(blog)
			h.Log.Debug("addBlog success: ", res.Success)
			h.Log.Debug("addBlog code: ", res.Code)
			http.Redirect(w, r, indexRt, http.StatusFound)

		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// GetBlog GetBlog
func (h *MCHandler) GetBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetBlog")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in GetBlog", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {
			bvars := mux.Vars(r)
			bidstr := bvars["bid"]
			bid, _ := strconv.ParseInt(bidstr, 10, 64)
			h.Log.Debug("blog id in edit", bid)
			var wg sync.WaitGroup
			bg := h.Delegate.GetBlog(bid)

			var sd SiteData
			sd.Canonical = template.URL(h.SiteURL + "/viewPost/" + bidstr)
			sd.OgURL = template.URL(h.SiteURL + "/viewPost/" + bidstr)
			sd.OgSiteName = h.SiteName
			sd.OgType = "article"
			//sd.Description = bg.Name
			sd.OgTitle = bg.Name

			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			bp.SiteData = &sd
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
				h.Log.Debug("TextHTML: ", bb.TextHTML)
			}

			bp.Blog = &bb
			wg.Add(1)
			go func(bbb *Blog) {
				defer wg.Done()
				cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
				h.Log.Debug("commentCnt: ", len(*cl))
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
				defer wg.Done()
				u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
				h.Log.Debug("get user: ")
				bbb.User = u1
				bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
				h.Log.Debug("User Done: ")
			}(&bb)

			wg.Wait()

			h.Templates.ExecuteTemplate(w, blogPage, &bp)

		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// UpdateBlogPage UpdateBlogPage
func (h *MCHandler) UpdateBlogPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in UpdateBlogPage")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in UpdateBlogPage", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in GetBlog: ", loggedInAuth)
		if loggedInAuth == true {
			bvars := mux.Vars(r)
			bidstr := bvars["bid"]
			bid, _ := strconv.ParseInt(bidstr, 10, 64)
			h.Log.Debug("blog id in edit", bid)
			bg := h.Delegate.GetBlog(bid)

			var ubp BlogPage
			ubp.Title = h.Title
			ubp.Desc = h.Desc
			ubp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				ubp.MyEmail = uemail.(string)
			}
			var ubb Blog
			ubb.Blog = bg
			txt, err := b64.StdEncoding.DecodeString(bg.Content)
			if err == nil {
				bg.Content = string(txt)
				bg.Content = strings.Replace(bg.Content, stripOut, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut2, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut3, "", -1)
				bg.Content = strings.Replace(bg.Content, stripOut4, "", -1)

				ubb.TextHTML = template.HTML(bg.Content)
				h.Log.Debug("TextHTML: ", ubb.TextHTML)
			}

			ubp.Blog = &ubb

			u1 := h.Delegate.GetUserByID(ubb.Blog.UserID)
			h.Log.Debug("get user: ")

			ubb.User = u1
			ubb.UserImage = b64.StdEncoding.EncodeToString(ubb.User.Image)
			h.Log.Debug("User Done: ")

			if ubp.Blog.User.Email == ubp.MyEmail {
				h.Templates.ExecuteTemplate(w, editBlogPage, &ubp)
			} else {
				http.Redirect(w, r, indexRt, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

// UpdateBlog UpdateBlog
func (h *MCHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in UpdateBlog")
	idStr := r.FormValue("id")
	h.Log.Debug("id in updateBlog: ", idStr)
	name := r.FormValue("name")
	h.Log.Debug("name in addBlog: ", name)
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in addBlog: ", loggedInAuth)
		if loggedInAuth == true {
			_, blog := h.processBlog(r)
			h.Log.Debug("UserID in updateBlog: ", blog.UserID)
			h.Log.Debug("id in updateBlog: ", blog.ID)
			h.Log.Debug("name in addBlog name: ", blog.Name)
			h.Log.Debug("blog in addBlog content: ", blog.Content)
			res := h.Delegate.UpdateBlog(blog)
			h.Log.Debug("addBlog success: ", res.Success)
			h.Log.Debug("addBlog code: ", res.Code)
			http.Redirect(w, r, indexRt, http.StatusFound)

		} else {
			http.Redirect(w, r, indexRt, http.StatusFound)
		}
	}
}

// SearchBlogList SearchBlogList
func (h *MCHandler) SearchBlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in SearchBlogList")
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in SearchBlogList", suc)
	if suc {
		loggedInAuth := s.Get("loggedIn")
		h.Log.Debug("loggedIn in SearchBlogList: ", loggedInAuth)
		if loggedInAuth == true {
			name := r.FormValue("name")
			sblogList := h.Delegate.GetBlogByName(name, 0, maxPosts)
			h.Log.Debug("blogCnt: ", len(*sblogList))
			var sbp BlogPage
			sbp.Title = h.Title
			sbp.Desc = h.Desc
			sbp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				sbp.MyEmail = uemail.(string)
			}
			var isAdmin = s.Get("isAdmin")
			h.Log.Debug("isAdmin: ", isAdmin)
			if isAdmin == true {
				sbp.IsAdmin = true
			}
			var blst []Blog
			for i := range *sblogList {
				var wg sync.WaitGroup
				var sbb Blog
				sbb.Blog = &(*sblogList)[i]
				txt, err := b64.StdEncoding.DecodeString(sbb.Blog.Content)
				if err == nil {
					sbb.Blog.Content = string(txt)
					sbb.Blog.Content = strings.Replace(sbb.Blog.Content, stripOut, "", -1)
					sbb.Blog.Content = strings.Replace(sbb.Blog.Content, stripOut2, "", -1)
					sbb.Blog.Content = strings.Replace(sbb.Blog.Content, stripOut3, "", -1)
					sbb.Blog.Content = strings.Replace(sbb.Blog.Content, stripOut4, "", -1)

					sbb.TextHTML = template.HTML(sbb.Blog.Content)
					h.Log.Debug("TextHTML: ", sbb.TextHTML)
				}

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, maxComments)
					h.Log.Debug("commentCnt: ", len(*cl))
					bbb.CommentCnt = len(*cl)
					h.Log.Debug("Comments Done: ")
				}(&sbb)
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
				}(&sbb)

				wg.Add(1)
				go func(bbb *Blog) {
					defer wg.Done()
					u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
					h.Log.Debug("get user: ")
					bbb.User = u1
					bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)
					h.Log.Debug("User Done: ")
				}(&sbb)

				wg.Wait()

				blst = append(blst, sbb)
			}
			sbp.BlogList = &blst
			h.Log.Debug("after all waits")

			h.Templates.ExecuteTemplate(w, blogListPage, &sbp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processBlog(r *http.Request) (string, *mcd.Blog) {
	var rtn mcd.Blog
	idStr := r.FormValue("id")
	uidStr := r.FormValue("userId")
	name := r.FormValue("name")
	email := r.FormValue("userEmail")
	content := r.FormValue("content")
	h.Log.Debug("name in processBlog: ", r.FormValue("name"))
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		rtn.ID = id
	}
	if uidStr != "" {
		uid, _ := strconv.ParseInt(uidStr, 10, 64)
		rtn.UserID = uid
	}
	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	rtn.Name = name
	rtn.Entered = time.Now()

	return email, &rtn
}
