package confighub

type Boolean struct {
	Value bool
}

func (b Boolean) String() (string, error) {
	return "", TypeValueErr
}

func (b Boolean) Boolean() (bool, error) {
	return b.Value, nil
}

func (b Boolean) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b Boolean) Long() (int64, error) {
	return 0, TypeValueErr
}

func (b Boolean) Float() (float32, error) {
	return 0, TypeValueErr
}

func (b Boolean) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b Boolean) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b Boolean) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}

func (b Boolean) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *Boolean) Parse(value interface{}) (err error) {

	valS, ok := value.(string)
	if !ok {
		return WrongValueErr
	}

	switch valS {
	case "true":
		b.Value = true
	case "false":
		b.Value = false
	default:
		err = WrongValueErr
	}

	return
}
