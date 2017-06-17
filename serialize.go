package config

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func encodeGob(object interface{}) ([]byte, error) {
	writer := bytes.NewBufferString("")
	encoder := gob.NewEncoder(writer)
	err := encoder.Encode(object)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func decodeGob(input []byte, object interface{}) error {
	writer := bytes.NewBuffer(input)
	decoder := gob.NewDecoder(writer)
	err := decoder.Decode(object)
	if err != nil {
		return err
	}
	return nil
}

func encodeJSON(object interface{}) ([]byte, error) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func decodeJSON(input []byte, object interface{}) error {
	return json.Unmarshal(input, object)
}
