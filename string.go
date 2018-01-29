package confighub

type String struct {
	Value string
}

func (b String) String() (string, error) {
	return b.Value, nil
}

func (b String) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b String) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b String) Long() (int64, error) {
	return 0, TypeValueErr
}

func (b String) Float() (float32, error) {
	return 0, TypeValueErr
}

func (b String) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b String) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}

func (b String) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b String) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *String) Parse(value interface{}) (err error) {

	valS, ok := value.(string)
	if !ok {
		return WrongValueErr
	}

	b.Value = valS
	return
}
