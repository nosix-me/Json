// +build !go1.1

// Copyright 2013 Json authors
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
	"log"
	"strconv"
	"testing"
)

func TestJsonGo10(t *testing.T) {
	js, err := NewJson([]byte(`{ 
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 8000000000
		}
	}`))
	if err != nil {
		log.Println("NewJson() error!")
	}
	log.Println(nil == js)

	arr, _ := js.Get("test").Get("array").Array()
	log.Println(nil == arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case float64:
			iv = int(v.(float64))
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		log.Println(i+1 == iv)
	}

	ma := js.Get("test").Get("array").MustArray()
	log.Println(ma == []interface{}{float64(1), "2", float64(3)})

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	log.Println(mm == map[string]interface{}{"subkeyone": float64(1)})

	log.Println(js.Get("test").Get("bignum").MustInt64() == int64(8000000000))
}
