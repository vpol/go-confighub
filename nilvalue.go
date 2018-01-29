package confighub

type NilValue struct {}

func (b NilValue) String() (string, error) {
	return "", TypeValueErr
}

func (b NilValue) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b NilValue) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b NilValue) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b NilValue) Float() (float32, error) {
	return 0, TypeValueErr
}
func (b NilValue) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b NilValue) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b NilValue) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}

func (b NilValue) File() (*File, error) {
	return nil, TypeValueErr
}

func (b *NilValue) Parse(value interface{}) (err error) {
	return nil
}

