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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("hello", "world")

	str := "{\"hello\":\"world\"}"
	body := NewBody(str)

	msg := New(header, body)

	assert.NotNil(msg)
	assert.NotNil(msg.Header)
	assert.NotNil(msg.Body)

	assert.True(msg.Header.Exist("hello"))

	json, err := msg.Body.JSON()
	assert.Nil(err)
	assert.Equal(json, str)
}

func TestMessage_JSON(t *testing.T) {
	assert := assert.New(t)

	strBody := "{\"hello\":\"world\"}"
	strHeader := "{\"hello\":\"world\"}"

	header := Header{}
	header.Add("hello", "world")

	body := NewBody(strBody)

	msg := New(header, body)
	str, err := msg.JSON()

	assert.Nil(err)
	assert.NotEmpty(str)
	assert.Equal(str, fmt.Sprintf("{\"header\":%s,\"body\":%s,\"status\":%d}", strHeader, strBody, 0))
}
