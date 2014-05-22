// +build !go1.2

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
	"log"
	"strconv"
	"testing"
)

func TestJsonGo12(t *testing.T) {
	log.Println("TestJsonGo12")
	js, err := NewJson([]byte(`{ 
		"test": { 
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807
		}
	}`))

	log.Println(nil == js)
	log.Println(nil == err)

	arr, _ := js.Get("test").Get("array").Array()
	log.Println(nil == arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			log.Println(nil == err)
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		log.Println(i+1 == iv)
	}

	ma := js.Get("test").Get("array").MustArray()
	log.Println(ma)

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	log.Println(mm)

	log.Println(js.Get("test").Get("bignum").MustInt64() == int64(9223372036854775807))
}
