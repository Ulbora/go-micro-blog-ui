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
}

// AddBlogPageData AddBlogPageData
type AddBlogPageData struct {
	Title     string
	Desc      string
	KeyWords  string
	UserEmail string
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
			blogList := h.Delegate.GetBlogList(0, 100)
			h.Log.Debug("blogCnt: ", len(*blogList))
			var bp BlogPage
			bp.Title = h.Title
			bp.Desc = h.Desc
			bp.KeyWords = h.KeyWords
			uemail := s.Get("userEmail")
			if uemail != nil {
				bp.MyEmail = uemail.(string)
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
					h.Log.Debug("TextHTML: ", bb.TextHTML)
				}
				//var ccnt int
				//var lcnt int
				wg.Add(1)
				go func(bbb *Blog) {
					// wg.Add(1)
					defer wg.Done()
					cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, 1000)
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
			//h.Log.Debug("template: ", h.AdminTemplates)
			//res := h.Service.GetContentList(false)
			// sort.Slice(*res, func(p, q int) bool {
			// 	return (*res)[p].Title < (*res)[q].Title
			// })
			// h.Log.Debug("content in admin index sorted: ", *res)
			h.Templates.ExecuteTemplate(w, blogListPage, &bp)
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
	//<object data="data:application/pdf;base64,YOURBASE64DATA" type="application/pdf"></object>
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

			//res := h.Service.GetContentList(false)
			// sort.Slice(*res, func(p, q int) bool {
			// 	return (*res)[p].Title < (*res)[q].Title
			// })
			// h.Log.Debug("content in admin index sorted: ", *res)
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
				cl := h.Delegate.GetCommentList(bbb.Blog.ID, 0, 1000)
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
			//var wg sync.WaitGroup
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

			u1 := h.Delegate.GetUserByID(bb.Blog.UserID)
			h.Log.Debug("get user: ")

			bb.User = u1
			bb.UserImage = b64.StdEncoding.EncodeToString(bb.User.Image)
			h.Log.Debug("User Done: ")

			if bp.Blog.User.Email == bp.MyEmail {
				h.Templates.ExecuteTemplate(w, editBlogPage, &bp)
			} else {
				http.Redirect(w, r, indexRt, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, loginRt, http.StatusFound)
		}
	}
}

func (h *MCHandler) processBlog(r *http.Request) (string, *mcd.Blog) {
	var rtn mcd.Blog
	name := r.FormValue("name")
	email := r.FormValue("userEmail")
	content := r.FormValue("content")
	h.Log.Debug("name in processBlog: ", r.FormValue("name"))

	rtn.Content = b64.StdEncoding.EncodeToString([]byte(content))
	rtn.Name = name
	rtn.Entered = time.Now()

	return email, &rtn
}
