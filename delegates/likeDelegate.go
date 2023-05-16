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

// AddLike AddLike
func (d *MCDelegate) AddLike(l *Like) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(l)
	if err == nil {
		lrq, err := d.buildRequest(http.MethodPost, "/rs/like/add", aJSON, apiKey)
		if err == nil {
			suc, stat := d.proxy.Do(lrq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// RemoveLike RemoveLike
func (d *MCDelegate) RemoveLike(l *Like) *Response {
	var rtn Response
	bidStr := strconv.FormatInt(l.BlogID, 10)
	uidStr := strconv.FormatInt(l.UserID, 10)
	// aJSON, err := json.Marshal(l)
	// if err == nil {
	dlrq, err := d.buildRequest(http.MethodDelete, "/rs/like/remove/"+bidStr+"/"+uidStr, nil, apiKey)
	if err == nil {
		usuc, stat := d.proxy.Do(dlrq, &rtn)
		rtn.Code = int64(stat)
		d.Log.Debug("suc: ", usuc)
		d.Log.Debug("stat: ", stat)
	}
	// }
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}


// ViewLikes ViewLikes
func (d *MCDelegate) ViewLikes(bid int64) *[]Like {
	var rtn []Like
	bidStr := strconv.FormatInt(bid, 10)
	// endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/like/view/"+bidStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
