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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader_Add(t *testing.T) {
	assert := assert.New(t)

	//one element test
	header := Header{}
	header.Add("001", "test-001")

	assert.Len(header, 1)
	assert.Contains(header, "001")
}

func TestHeader_Add1000(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	for i := 0; i < 1000; i++ {
		header.Add(strconv.Itoa(i), "some test")
	}

	assert.Len(header, 1000)
	for i := 0; i < 999; i++ {
		assert.Contains(header, strconv.Itoa(i))
	}
}

func TestHeader_Remove(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	assert.Len(header, 1)
	header.Remove("1")
	assert.Len(header, 0)
}

func TestHeader_RemoveNone(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	assert.NotEmpty(header)
	assert.Len(header, 1)
	header.Remove("0")
	assert.Len(header, 1)
}

func TestHeader_RemoveOne(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	for i := 0; i < 1000; i++ {
		header.Add(strconv.Itoa(i), "some test")
	}

	assert.Len(header, 1000)
	header.Remove(strconv.Itoa(0))
	assert.Len(header, 999)
}

func TestHeader_RemoveAll(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	for i := 0; i < 1000; i++ {
		header.Add(strconv.Itoa(i), "some test")
	}

	assert.Len(header, 1000)
	for i := 0; i < 1000; i++ {
		header.Remove(strconv.Itoa(i))
	}
	assert.Empty(header)
}

func TestHeader_Get(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	value := header.Get("1")
	assert.NotEmpty(value)
	assert.Equal(value, "some test")
	assert.NotEmpty(header)
}

func TestHeader_GetNone(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	value := header.Get("0")
	assert.Empty(value)
	assert.NotEmpty(header)
}

func TestHeader_Get1000(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	for i := 0; i < 1000; i++ {
		header.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
	}

	assert.Len(header, 1000)
	for i := 0; i < 1000; i++ {
		value := header.Get(strconv.Itoa(i))
		assert.Equal(value, fmt.Sprintf("Some Test %d", i))
	}

	assert.Len(header, 1000)
}

func TestHeader_Exist(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	value := header.Exist("1")
	assert.True(value)
	assert.NotEmpty(header)
	assert.Len(header, 1)
}

func TestHeader_ExistNone(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	header.Add("1", "some test")

	value := header.Exist("0")
	assert.False(value)
	assert.NotEmpty(header)
	assert.Len(header, 1)
}

func TestHeader_Exist1000(t *testing.T) {
	assert := assert.New(t)

	header := Header{}
	for i := 0; i < 1000; i++ {
		header.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
	}

	assert.Len(header, 1000)
	for i := 0; i < 1000; i++ {
		value := header.Exist(strconv.Itoa(i))
		assert.True(value)
	}

	assert.Len(header, 1000)
}

func TestHeader_Equals(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{}
	header2 := Header{}
	for i := 0; i < 1000; i++ {
		header1.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
		header2.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
	}

	assert.True(header1.Equals(header2))
}

func TestHeader_EqualsInverse(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{}
	header2 := Header{}
	for i := 0; i < 1000; i++ {
		header1.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
	}

	for i := 999; i >= 0; i-- {
		header2.Add(strconv.Itoa(i), fmt.Sprintf("Some Test %d", i))
	}

	assert.True(header1.Equals(header2))
}

func TestHeader_EqualsError(t *testing.T) {
	assert := assert.New(t)

	header1 := Header{}
	header2 := Header{}

	header1.Add("hello", "world")
	header2.Add("hello", "le monde")

	assert.False(header1.Equals(header2))
}
