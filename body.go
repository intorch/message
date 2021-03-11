// Copyright 2021 Intorch (intorch.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

import (
	"encoding/json"
	"reflect"
)

//Body map to store JSON payload
type Body map[string]interface{}

//NewBody function to create a body object from JSON string
func NewBody(str string) Body {
	var body Body

	err := json.Unmarshal([]byte(str), &body)

	if err != nil {
		return nil
	}

	return body
}

//JSON function that parser this object in a JSON string
func (body Body) JSON() (string, error) {
	js, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	return string(js), nil
}

//Equals check if header is equals another one
func (body Body) Equals(other Body) bool {
	if reflect.DeepEqual(body, other) {
		return true
	}

	return false
}
