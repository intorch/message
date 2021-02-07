package message

import "encoding/json"

//Message data structure to store the transactional message
//object;
type Message struct {
	Header *Header `json:"header"`
	Body   *Body   `json:"body"`
	Status int     `json:"status"`
}

//NewMessage function to create a new Message structure
func NewMessage(header Header, body Body) *Message {
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
