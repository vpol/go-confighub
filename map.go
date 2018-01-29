package confighub

type StringMap struct {
	Value map[string]string
}

func (b StringMap) String() (string, error) {
	return "", TypeValueErr
}

func (b StringMap) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b StringMap) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b StringMap) Long() (int64, error) {
	return 0, TypeValueErr
}

func (b StringMap) StringMap() (map[string]string, error) {
	return b.Value, nil
}

func (b StringMap) Float() (float32, error) {
	return 0, TypeValueErr
}

func (b StringMap) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b StringMap) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b StringMap) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *StringMap) Parse(value interface{}) (err error) {

	valM, ok := value.(map[string]interface{})
	if !ok {
		return WrongValueErr
	}

	b.Value = make(map[string]string)

	for k, v := range valM {
		b.Value[k] = v.(string)
	}

	return
}
