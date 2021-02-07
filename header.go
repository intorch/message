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
