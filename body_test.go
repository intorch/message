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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody_JSON(t *testing.T) {
	assert := assert.New(t)
	body := Body{}

	body["hello"] = "world"
	js, err := body.JSON()

	assert.Nil(err)
	assert.NotEmpty(js)
	assert.Equal(js, "{\"hello\":\"world\"}")
}

func TestNewBody(t *testing.T) {
	assert := assert.New(t)

	body := NewBody("{\"hello\": \"world\"}")

	assert.NotNil(body)
	assert.NotEmpty(body)
	assert.Len(body, 1)
	assert.Contains(body, "hello")
	assert.Equal(body["hello"].(string), "world")
}

func TestBody_Equals(t *testing.T) {
	assert := assert.New(t)

	body1 := NewBody("{\"hello\": \"world\"}")
	body2 := NewBody("{\"hello\": \"world\"}")

	assert.True(body1.Equals(body2))
}

func TestBody_EqualsDeep(t *testing.T) {
	assert := assert.New(t)

	body1 := NewBody("{\"hello\": {\"hello\": \"world\"}}")
	body2 := NewBody("{\"hello\": {\"hello\": \"le monde\"}}")

	assert.False(body1.Equals(body2))
}

func TestBody_EqualsFalse(t *testing.T) {
	assert := assert.New(t)

	body1 := NewBody("{\"hello\": \"world\"}")
	body2 := NewBody("{\"hello\": \"le monde\"}")

	assert.False(body1.Equals(body2))
}
