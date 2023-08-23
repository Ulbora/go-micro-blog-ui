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

// AddUser AddUser
func (d *MCDelegate) AddUser(u *User) *ResponseID {
	var rtn ResponseID
	aJSON, err := json.Marshal(u)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPost, "/rs/user/add", aJSON, apiKey)
		if err == nil {
			suc, stat := d.proxy.Do(rq, &rtn)
			rtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rtn)
	return &rtn
}

// UpdateUser UpdateUser
func (d *MCDelegate) UpdateUser(u *User) *Response {
	var rtn Response
	aJSON, err := json.Marshal(u)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/user/update", aJSON, apiKey)
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

// GetUser GetUser
func (d *MCDelegate) GetUser(email string) *User {
	var rtn User
	rq, err := d.buildRequest(http.MethodGet, "/rs/user/get/"+email, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetUserByID GetUserByID
func (d *MCDelegate) GetUserByID(id int64) *User {
	var rtn User
	uidStr := strconv.FormatInt(id, 10)
	rq, err := d.buildRequest(http.MethodGet, "/rs/user/get/id/"+uidStr, nil, apiKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetUserList GetUserList
func (d *MCDelegate) GetUserList() *[]User {
	var rtn []User
	rq, err := d.buildRequest(http.MethodGet, "/rs/user/list", nil, adminKey)
	if err == nil {
		lsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", lsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetUnActivatedUserList GetUnActivatedUserList
func (d *MCDelegate) GetUnActivatedUserList() *[]User {
	var rtn []User
	rq, err := d.buildRequest(http.MethodGet, "/rs/user/unactivivated/list", nil, adminKey)
	if err == nil {
		lsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", lsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// GetBannedUserList GetBannedUserList
func (d *MCDelegate) GetBannedUserList() *[]User {
	var rtn []User
	rq, err := d.buildRequest(http.MethodGet, "/rs/user/banned/list", nil, adminKey)
	if err == nil {
		lsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", lsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// EnableUser EnableUser
func (d *MCDelegate) EnableUser(u *User) *Response {
	var eurtn Response
	aJSON, err := json.Marshal(u)
	if err == nil {
		erq, err := d.buildRequest(http.MethodPut, "/rs/user/enable", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(erq, &eurtn)
			eurtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", eurtn)
	return &eurtn
}

// DisableUser DisableUser
func (d *MCDelegate) DisableUser(u *User) *Response {
	var durtn Response
	aJSON, err := json.Marshal(u)
	if err == nil {
		drq, err := d.buildRequest(http.MethodPut, "/rs/user/disable", aJSON, adminKey)
		if err == nil {
			dsuc, stat := d.proxy.Do(drq, &durtn)
			durtn.Code = int64(stat)
			d.Log.Debug("suc: ", dsuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", durtn)
	return &durtn
}

// DisableUserForCause DisableUserForCause
func (d *MCDelegate) DisableUserForCause(u *User) *Response {
	var ddurtn Response
	aJSON, err := json.Marshal(u)
	if err == nil {
		drq, err := d.buildRequest(http.MethodPut, "/rs/user/disableForCause", aJSON, adminKey)
		if err == nil {
			dsuc, stat := d.proxy.Do(drq, &ddurtn)
			ddurtn.Code = int64(stat)
			d.Log.Debug("suc: ", dsuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", ddurtn)
	return &ddurtn
}

// ReactivateUser ReactivateUser
func (d *MCDelegate) ReactivateUser(u *User) *Response {
	var rertn Response
	aJSON, err := json.Marshal(u)
	if err == nil {
		erq, err := d.buildRequest(http.MethodPut, "/rs/user/reinstate", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(erq, &rertn)
			rertn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", rertn)
	return &rertn
}
