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

// AddAbout AddAbout
func (d *MCDelegate) AddAbout(a *About) *ResponseID {
	var arrtn ResponseID
	aJSON, err := json.Marshal(a)
	if err == nil {
		aarq, err := d.buildRequest(http.MethodPost, "/rs/about/add", aJSON, adminKey)
		if err == nil {
			aasuc, stat := d.proxy.Do(aarq, &arrtn)
			arrtn.Code = int64(stat)
			d.Log.Debug("suc: ", aasuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", arrtn)
	return &arrtn
}

// UpdateAbout UpdateAbout
func (d *MCDelegate) UpdateAbout(r *About) *Response {
	var uartn Response
	aJSON, err := json.Marshal(r)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/about/update", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(rq, &uartn)
			uartn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", uartn)
	return &uartn
}

// GetAbout GetAbout
func (d *MCDelegate) GetAbout() *About {
	var rtn About
	arq, err := d.buildRequest(http.MethodGet, "/rs/about/get", nil, adminKey)
	if err == nil {
		gasuc, stat := d.proxy.Do(arq, &rtn)
		d.Log.Debug("suc: ", gasuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
