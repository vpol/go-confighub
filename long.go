package confighub

import (
	"strconv"
)

type Long struct {
	Value int64
}

func (b Long) String() (string, error) {
	return "", TypeValueErr
}

func (b Long) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b Long) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b Long) Long() (int64, error) {
	return b.Value, nil
}
func (b Long) Float() (float32, error) {
	return 0, TypeValueErr
}
func (b Long) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b Long) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b Long) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *Long) Parse(value string) (err error) {
	b.Value, err = strconv.ParseInt(value, 10, 0)
	return
}
