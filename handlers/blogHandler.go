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
	"net/http"
	"sync"

	b64 "encoding/base64"

	mcd "github.com/Ulbora/go-micro-blog-ui/delegates"
)

// Blog Blog
type Blog struct {
	Blog       *mcd.Blog
	User       *mcd.User
	UserImage  string
	CommentCnt int
	LikeCnt    int
}

// BlogPage BlogPage
type BlogPage struct {
	Title    string
	Desc     string
	KeyWords string
	BlogList *[]Blog
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
			var blst []Blog
			for i := range *blogList {
				var wg sync.WaitGroup
				var bb Blog
				bb.Blog = &(*blogList)[i]
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

				}(&bb)
				wg.Add(1)
				go func(bbb *Blog) {
					// wg.Add(1)
					defer wg.Done()
					ll := h.Delegate.ViewLikes(bbb.Blog.ID)
					h.Log.Debug("likeCnt: ", len(*ll))
					// bb.LikeCnt = len(*ll)
					bbb.LikeCnt = len(*ll)
				}(&bb)

				wg.Add(1)
				go func(bbb *Blog) {
					// wg.Add(1)
					defer wg.Done()
					u1 := h.Delegate.GetUserByID(bbb.Blog.UserID)
					//h.Log.Debug("commentCnt: ", )
					// bb.CommentCnt = len(*cl)
					//ccnt = len(*cl)
					bbb.User = u1
					bbb.UserImage = b64.StdEncoding.EncodeToString(bbb.User.Image)

				}(&bb)

				wg.Wait()
				h.Log.Debug("after wait: for ", i)
				// h.Log.Debug("comments : ", ccnt)
				// h.Log.Debug("likes : ", lcnt)
				// bb.CommentCnt = ccnt
				// bb.LikeCnt = lcnt

				blst = append(blst, bb)
				h.Log.Debug("blog: ", bb)
			}
			bp.BlogList = &blst
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
}
