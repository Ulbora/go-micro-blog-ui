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

// AddRule AddRule
func (d *MCDelegate) AddRule(r *Rule) *ResponseID {
	var arrtn ResponseID
	aJSON, err := json.Marshal(r)
	if err == nil {
		brq, err := d.buildRequest(http.MethodPost, "/rs/rule/add", aJSON, adminKey)
		if err == nil {
			suc, stat := d.proxy.Do(brq, &arrtn)
			arrtn.Code = int64(stat)
			d.Log.Debug("suc: ", suc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", arrtn)
	return &arrtn
}

// UpdateRule UpdateRule
func (d *MCDelegate) UpdateRule(r *Rule) *Response {
	var urrtn Response
	aJSON, err := json.Marshal(r)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/rule/update", aJSON, adminKey)
		if err == nil {
			usuc, stat := d.proxy.Do(rq, &urrtn)
			urrtn.Code = int64(stat)
			d.Log.Debug("suc: ", usuc)
			d.Log.Debug("stat: ", stat)
		}
	}
	d.Log.Debug("rtn: ", urrtn)
	return &urrtn
}

// GetRule GetRule
func (d *MCDelegate) GetRule() *Rule {
	var rtn Rule
	brq, err := d.buildRequest(http.MethodGet, "/rs/rule/get", nil, adminKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(brq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}
