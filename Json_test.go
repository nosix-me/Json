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
	"testing"
)

func TestJson(t *testing.T) {
	js, err := NewJson([]byte(`{ 
		"test": { 
			"string_array": ["asdf", "ghjk", "zxcv"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "Json",
			"bool": true 
		}
	}`))
	if err != nil {
		log.Panicln("NewJson() error!")
	}
	aws := js.Get("test").Get("array")
	log.Println(aws)

	_, ok := js.CheckGet("test")
	log.Println(ok)

	_, ok = js.CheckGet("missing_key")
	log.Println(ok)

	aws = js.Get("test").Get("arraywithsubs")
	log.Println(aws)

	var awsval int

	awsval, _ = aws.GetIndex(0).Get("subkeyone").Int()
	log.Println(awsval)

	awsval, _ = aws.GetIndex(1).Get("subkeytwo").Int()
	log.Println(awsval)

	awsval, _ = aws.GetIndex(1).Get("subkeythree").Int()
	log.Println(awsval)

	i, _ := js.Get("test").Get("int").Int()
	log.Println(i)

	f, _ := js.Get("test").Get("float").Float64()
	log.Println(f)

	s, _ := js.Get("test").Get("string").String()
	log.Println(s)

	b, _ := js.Get("test").Get("bool").Bool()
	log.Println(b)

	mi := js.Get("test").Get("int").MustInt()
	log.Println(mi)

	mi2 := js.Get("test").Get("missing_int").MustInt(5150)
	log.Println(mi2)

	ms := js.Get("test").Get("string").MustString()
	log.Println(ms)

	ms2 := js.Get("test").Get("missing_string").MustString("fyea")
	log.Println(ms2)

	ma2 := js.Get("test").Get("missing_array").MustArray([]interface{}{"1", 2, "3"})
	log.Println(ma2)

	mm2 := js.Get("test").Get("missing_map").MustMap(map[string]interface{}{"found": false})
	log.Println(mm2)

	strs, err := js.Get("test").Get("string_array").StringArray()
	if err != nil {
		log.Println(`js.Get("test").Get("string_array").StringArray() error`)
	}
	log.Println(strs[0])
	log.Println(strs[1])
	log.Println(strs[2])

	gp, _ := js.GetPath("test", "string").String()
	log.Println(gp)

	gp2, _ := js.GetPath("test", "int").Int()
	log.Println(gp2)

	log.Println(js.Get("test").Get("bool").MustBool() == true)

	js.Set("float2", 300.0)
	log.Println(js.Get("float2").MustFloat64() == 300.0)

	js.Set("test2", "setTest")
	log.Println("setTest" == js.Get("test2").MustString())
}

func TestStdlibInterfaces(t *testing.T) {
	val := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})
	val2 := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})

	raw := `{"name":"myobject","params":{"string":"Json"}}`
	json.Unmarshal([]byte(raw), val)

	log.Println("myobject" == val.Name)

	s, _ := val.Params.Get("string").String()
	log.Println("Json" == s)

	p, err := json.Marshal(val)
	if err != nil {
		log.Println(`json.Marshal(val) error!`)
	}
	log.Println(json.Unmarshal(p, val2) == nil)
	log.Println(val == val2) // stable
}
