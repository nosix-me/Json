// +build !go1.1

// Copyright 2014 Json authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Json is a package for json resolve
package Json

import (
	"encoding/json"
	"errors"
)

func (j *Json) UnmarshalJSON(p []byte) error {
	return json.Unmarshal(p, &j.data)
}

func (j *Json) Float64() (float64, error) {
	if i, ok := (j.data).(float64); ok {
		return i, nil
	}
	return -1, errors.New("Type assertion to float64 failed")
}

func (j *Json) Int() (int, error) {
	if f, ok := (j.data).(float64); ok {
		return int(f), nil
	}
	return -1, errors.New("Type assertion to float64 failed")
}

func (j *Json) Int64() (int64, error) {
	if f, ok := (j.data).(float64); ok {
		return int64(f), nil
	}
	return -1, errors.New("Type assertion to float64 failed")
}
