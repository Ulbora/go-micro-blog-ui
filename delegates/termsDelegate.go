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

// AddTerms AddTerms
func (d *MCDelegate) AddTerms(t *Terms) *ResponseID {
	var atrtn ResponseID
	aJSON, err := json.Marshal(t)
	if err == nil {
		atrq, err := d.buildRequest(http.MethodPost, "/rs/tos/add", aJSON, adminKey)
		if err == nil {
			suc, stat := d.proxy.Do(atrq, &atrtn)
			atrtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", atrtn)
	return &atrtn
}

// UpdateTerms UpdateTerms
func (d *MCDelegate) UpdateTerms(t *Terms) *Response {
	var utrtn Response
	aJSON, err := json.Marshal(t)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/tos/update", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(rq, &utrtn)
			utrtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", utrtn)
	return &utrtn
}

// GetTerms GetTerms
func (d *MCDelegate) GetTerms() *Terms {
	var rtn Terms
	trq, err := d.buildRequest(http.MethodGet, "/rs/tos/get", nil, adminKey)
	if err == nil {
		gtsuc, stat := d.proxy.Do(trq, &rtn)
		d.Log.Debug("suc: ", gtsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
