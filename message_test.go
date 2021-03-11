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

func TestMessage_EqualsFalseHeader1Nil(t *testing.T) {
	assert := assert.New(t)

	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(nil, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseHeader2Nil(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(nil, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseHeader1(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "le monde"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseHeader2(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "le monde"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseBody1Nil(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	msg1 := New(header1, nil)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseBody2Nil(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	msg2 := New(header2, nil)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseBody1(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"le monde\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseBody2(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"le mond\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsTrueStatusNotSet(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.True(msg1.Equals(msg2))
}

func TestMessage_EqualsTrueStatus(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)
	msg1.Status = 1

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)
	msg2.Status = 1

	assert.True(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseStatus2NotSet(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)
	msg1.Status = 1

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseStatus1NotSet(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)
	msg2.Status = 1

	assert.False(msg1.Equals(msg2))
}

func TestMessage_EqualsFalseStatus(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{"hello": "world"}
	body1 := NewBody("{\"hello\":\"world\"}")
	msg1 := New(header1, body1)
	msg1.Status = 1

	header2 := Header{"hello": "world"}
	body2 := NewBody("{\"hello\":\"world\"}")
	msg2 := New(header2, body2)
	msg2.Status = 2

	assert.False(msg1.Equals(msg2))
}
