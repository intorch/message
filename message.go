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
)

//Message data structure to store the transactional message
//object;
type Message struct {
	Header *Header `json:"header"`
	Body   *Body   `json:"body"`
	Status int     `json:"status"`
}

//New function to create a new Message structure
func New(header Header, body Body) *Message {
	return &Message{
		Header: &header,
		Body:   &body,
	}
}

//JSON function to  transform message in a JSON
func (msg *Message) JSON() (string, error) {
	bmsg, err := json.Marshal(msg)

	if err != nil {
		return "", err
	}

	return string(bmsg), nil
}

//Equals check if message is equals another one
func (msg Message) Equals(other *Message) bool {
	if other == nil {
		return false
	}

	if msg.Body != nil {
		if !msg.Body.Equals(*other.Body) {
			return false
		}
	} else {
		if other.Body != nil {
			return false
		}
	}

	if msg.Header != nil {
		if !msg.Header.Equals(*other.Header) {
			return false
		}
	} else {
		if other.Header != nil {
			return false
		}
	}

	if msg.Status != other.Status {
		return false
	}

	return true
}
