package confighub

type Value interface {
	Parse(interface{}) error
	Boolean() (bool, error)
	Integer() (int, error)
	Long() (int64, error)
	Float() (float32, error)
	Double() (float64, error)
	String() (string, error)
	File() (*File, error)
	StringMap() (map[string]string, error)
	Json() (map[string]interface{}, error)
}
