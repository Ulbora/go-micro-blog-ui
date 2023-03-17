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

// AddRole AddRole
func (d *MCDelegate) AddRole(r *Role) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(r)
	if err == nil {
		rrq, err := d.buildRequest(http.MethodPost, "/rs/role/add", aJSON, adminKey)
		if err == nil {
			suc, stat := d.proxy.Do(rrq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// GetRole GetRole
func (d *MCDelegate) GetRole(name string) *Role {
	var rtn Role
	rq, err := d.buildRequest(http.MethodGet, "/rs/role/get/"+name, nil, apiKey)
	if err == nil {
		grsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", grsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetRoleList GetRoleList
func (d *MCDelegate) GetRoleList() *[]Role {
	var rtn []Role
	rq, err := d.buildRequest(http.MethodGet, "/rs/role/list", nil, apiKey)
	if err == nil {
		lrsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", lrsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// DeleteRole DeleteRole
func (d *MCDelegate) DeleteRole(rid int64) *Response {
	var rtn Response
	ridStr := strconv.FormatInt(rid, 10)
	rq, err := d.buildRequest(http.MethodDelete, "/rs/role/delete/"+ridStr, nil, adminKey)
	if err == nil {
		grsuc, stat := d.proxy.Do(rq, &rtn)
		rtn.Code = int64(stat)
		d.Log.Debug("suc: ", grsuc)
		d.Log.Debug("stat: ", stat)
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}
