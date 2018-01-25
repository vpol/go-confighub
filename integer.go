package confighub

import (
	"strconv"
)

type Integer struct {
	Value int
}

func (b Integer) String() (string, error) {
	return "", TypeValueErr
}

func (b Integer) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b Integer) Integer() (int, error) {
	return b.Value, nil
}

func (b Integer) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b Integer) Float() (float32, error) {
	return 0, TypeValueErr
}
func (b Integer) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b Integer) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b Integer) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *Integer) Parse(value string) (err error) {
	var tVal int64
	tVal, err = strconv.ParseInt(value, 10, 32)
	if err != nil {
		return
	}
	b.Value = int(tVal)
	return
}
