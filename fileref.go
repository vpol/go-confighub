package confighub

type FileRef struct {
	Value   *File
	files *Files
}

func (b FileRef) String() (string, error) {
	return "", TypeValueErr
}

func (b FileRef) Boolean() (bool, error) {
	return false, TypeValueErr
}

func (b FileRef) Integer() (int, error) {
	return 0, TypeValueErr
}

func (b FileRef) Long() (int64, error) {
	return 0, TypeValueErr
}
func (b FileRef) Float() (float32, error) {
	return 0, TypeValueErr
}
func (b FileRef) Double() (float64, error) {
	return 0, TypeValueErr
}

func (b FileRef) Json() (map[string]interface{}, error) {
	return nil, TypeValueErr
}

func (b FileRef) File() (*File, error) {
	return b.Value, nil
}

func (b FileRef) StringMap() (map[string]string, error) {
	return nil, TypeValueErr
}


func (b *FileRef) Parse(value interface{}) (err error) {

	valS, ok := value.(string)
	if !ok {
		return WrongValueErr
	}

	if b.files != nil {
		if f, ok := b.files.entries[valS]; ok {
			b.Value = f
			return
		}
	}

	return FileRefNotFoundErr
}
