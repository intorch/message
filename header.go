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

//Header map to store some Message properties
type Header map[string]string

//Add create new key value pair. case the key already exist it'll be
//overwrited
func (header *Header) Add(key string, value string) {
	(*header)[key] = value
}

//Remove function that deletes the element with the specified key from
//the Header. If it is nil or there is no such element, Remove is a no-op.
func (header Header) Remove(key string) {
	delete(header, key)
}

//Get function to obtain the value from a key. If there is no such element
// it retuns empty string
func (header Header) Get(key string) string {
	return header[key]
}

//Exist function to check if there is the elemente mapped by key paramert
//inside the header object
func (header Header) Exist(key string) bool {
	if header[key] == "" {
		return false
	}

	return true
}
