package delegates

import (
	"encoding/json"
	"net/http"
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

// AddHome AddHome
func (d *MCDelegate) AddHome(a *Home) *ResponseID {
	var hrrtn ResponseID
	aJSON, err := json.Marshal(a)
	if err == nil {
		aarq, err := d.buildRequest(http.MethodPost, "/rs/home/add", aJSON, adminKey)
		if err == nil {
			aasuc, stat := d.proxy.Do(aarq, &hrrtn)
			hrrtn.Code = int64(stat)
			d.Log.Debug("suc: ", aasuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", hrrtn)
	return &hrrtn
}

// UpdateHome UpdateHome
func (d *MCDelegate) UpdateHome(r *Home) *Response {
	var uhrtn Response
	aJSON, err := json.Marshal(r)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/home/update", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(rq, &uhrtn)
			uhrtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", uhrtn)
	return &uhrtn
}

// GetHome GetHome
func (d *MCDelegate) GetHome() *Home {
	var rtn Home
	arq, err := d.buildRequest(http.MethodGet, "/rs/home/get", nil, adminKey)
	if err == nil {
		gasuc, stat := d.proxy.Do(arq, &rtn)
		d.Log.Debug("suc: ", gasuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
