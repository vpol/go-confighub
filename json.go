package confighub

import (
	"encoding/json"
)

type Json struct {
	Value map[string]interface{}
}

func (b Json) String() (string, error) {
	return "", TypeValueErr
}

func (b Json) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b Json) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b Json) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b Json) Float() (float32, error) {
	return 0, TypeValueErr
}

func (b Json) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b Json) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}

func (b Json) Json() (map[string]interface{}, error) {
	if b.Value == nil {
		return nil, TypeValueErr
	}
	return b.Value, nil
}

func (b Json) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *Json) Parse(value interface{}) (err error) {

	valS, ok := value.(string)
	if !ok {
		return WrongValueErr
	}

	err = json.Unmarshal([]byte(valS), &b.Value)
	return
}
