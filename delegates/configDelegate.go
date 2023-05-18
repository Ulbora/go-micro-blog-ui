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

// GetConfig GetConfig
func (d *MCDelegate) GetConfig() *Config {
	var rtn Config
	rq, err := d.buildRequest(http.MethodGet, "/rs/config/get", nil, adminKey)
	if err == nil {
		gsuc, stat := d.proxy.Do(rq, &rtn)
		d.Log.Debug("suc: ", gsuc)
		d.Log.Debug("stat: ", stat)
	}
	return &rtn
}

// UpdateConfig UpdateConfig
func (d *MCDelegate) UpdateConfig(c *Config) *Response {
	var rtn Response
	aJSON, err := json.Marshal(c)
	if err == nil {
		rq, err := d.buildRequest(http.MethodPut, "/rs/config/update", aJSON, adminKey)
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
