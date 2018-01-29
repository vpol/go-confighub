package confighub

import (
	"strconv"
)

type Float struct {
	Value float32
}

func (b Float) String() (string, error) {
	return "", TypeValueErr
}

func (b Float) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b Float) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b Float) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b Float) Float() (float32, error) {
	return b.Value, nil
}
func (b Float) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b Float) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b Float) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}

func (b Float) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *Float) Parse(value interface{}) (err error) {
	valS, ok := value.(string)
	if !ok {
		return WrongValueErr
	}

	var tVal float64
	tVal, err = strconv.ParseFloat(valS, 10)
	if err != nil {
		return
	}
	b.Value = float32(tVal)
	return
}
