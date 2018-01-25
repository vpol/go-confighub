package confighub

import (
	"strconv"
)

type Double struct {
	Value float64
}

func (b Double) String() (string, error) {
	return "", TypeValueErr
}

func (b Double) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b Double) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b Double) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b Double) Float() (float32, error) {
	return 0, TypeValueErr
}
func (b Double) Double() (float64, error) {
	return b.Value, nil
}

func (b Double) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b Double) File() (*File, error) {
	return nil, TypeValueErr
}


func (b *Double) Parse(value string) (err error) {
	b.Value, err = strconv.ParseFloat(value, 10)
	return
}
