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

func (b *FileRef) Parse(value string) (err error) {

	if b.files != nil {
		if f, ok := b.files.entries[value]; ok {
			b.Value = f
			return
		}
	}

	return FileRefNotFoundErr
}
