package delegates

import (
	"encoding/json"
	"net/http"
	"strconv"
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

// AddBlog AddBlog
func (d *MCDelegate) AddBlog(b *Blog) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(b)
	if err == nil {
		brq, err := d.buildRequest(http.MethodPost, "/rs/blog/add", aJSON, apiKey)
		if err == nil {
			suc, stat := d.proxy.Do(brq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// UpdateBlog UpdateBlog
func (d *MCDelegate) UpdateBlog(b *Blog) *Response {
	var rtn Response
	aJSON, err := json.Marshal(b)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/blog/update", aJSON, apiKey)
		if err == nil {
			usuc, stat := d.proxy.Do(rq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// GetBlog GetBlog
func (d *MCDelegate) GetBlog(bid int64) *Blog {
	var rtn Blog
	bidStr := strconv.FormatInt(bid, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/blog/get/id/"+bidStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetBlogByName GetBlogByName
func (d *MCDelegate) GetBlogByName(name string, start int64, end int64) *[]Blog {
	var rtn []Blog
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/blog/get/name/"+name+"/"+startStr+"/"+endStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetBlogList GetBlogList
func (d *MCDelegate) GetBlogList(start int64, end int64) *[]Blog {
	var rtn []Blog
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/blog/list/"+startStr+"/"+endStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetAdminBlogList GetAdminBlogList
func (d *MCDelegate) GetAdminBlogList(start int64, end int64) *[]Blog {
	var rtn []Blog
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/blog/admin/list/"+startStr+"/"+endStr, nil, adminKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// ActivateBlog ActivateBlog
func (d *MCDelegate) ActivateBlog(b *Blog) *Response {
	var rtn Response
	aJSON, err := json.Marshal(b)
	if err == nil {
		erq, err := d.buildRequest(http.MethodPut, "/rs/blog/activate", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(erq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// DeActivateBlog DeActivateBlog
func (d *MCDelegate) DeActivateBlog(b *Blog) *Response {
	var rtn Response
	aJSON, err := json.Marshal(b)
	if err == nil {
		darq, err := d.buildRequest(http.MethodPut, "/rs/blog/deactivate", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(darq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}
