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

// AddComment AddComment
func (d *MCDelegate) AddComment(c *Comment) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(c)
	if err == nil {
		brq, err := d.buildRequest(http.MethodPost, "/rs/comment/add", aJSON, apiKey)
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

// UpdateComment UpdateComment
func (d *MCDelegate) UpdateComment(c *Comment) *Response {
	var rtn Response
	aJSON, err := json.Marshal(c)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/comment/update", aJSON, apiKey)
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

// GetCommentList GetCommentList
func (d *MCDelegate) GetCommentList(blogID int64, start int64, end int64) *[]Comment {
	var rtn []Comment
	bidStr := strconv.FormatInt(blogID, 10)
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/comment/list/"+bidStr+"/"+startStr+"/"+endStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetAdminCommentList GetAdminCommentList
func (d *MCDelegate) GetAdminCommentList(blogID int64, start int64, end int64) *[]Comment {
	var rtn []Comment
	bidStr := strconv.FormatInt(blogID, 10)
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/comment/admin/list/"+bidStr+"/"+startStr+"/"+endStr, nil, adminKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// ActivateComment ActivateComment
func (d *MCDelegate) ActivateComment(c *Comment) *Response {
	var rtn Response
	aJSON, err := json.Marshal(c)
	if err == nil {
		erq, err := d.buildRequest(http.MethodPut, "/rs/comment/activate", aJSON, adminKey)
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

// DeActivateComment DeActivateComment
func (d *MCDelegate) DeActivateComment(c *Comment) *Response {
	var rtn Response
	aJSON, err := json.Marshal(c)
	if err == nil {
		darq, err := d.buildRequest(http.MethodPut, "/rs/comment/deactivate", aJSON, adminKey)
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
