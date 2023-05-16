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

// AddUserAuth AddUserAuth
func (d *MCDelegate) AddUserAuth(a *UserAuth) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(a)
	if err == nil {
		brq, err := d.buildRequest(http.MethodPost, "/rs/user/auth/add", aJSON, apiKey)
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

// GetUserAuthList GetUserAuthList
func (d *MCDelegate) GetUserAuthList(userID int64, start int64, end int64) *[]UserAuth {
	var rtn []UserAuth
	uidStr := strconv.FormatInt(userID, 10)
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)
	brq, err := d.buildRequest(http.MethodGet, "/rs/user/auth/list/"+uidStr+"/"+startStr+"/"+endStr, nil, adminKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
