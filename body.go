package message

import "encoding/json"

//Body map to store JSON payload
type Body map[string]interface{}

//JSON function that parser this object in a JSON string
func (body Body) JSON() (string, error) {
	js, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	return string(js), nil
}
