package confighub

import "encoding/json"

type Files struct {
	cfg     *ConfigHubClient
	entries map[string]*File
}

func (f *Files) UnmarshalJSON(data []byte) (err error) {

	var m = make(map[string]struct {
		Content     string `json:"content"`
		ContentType string `json:"content-type"`
	})

	err = json.Unmarshal(data, &m)

	if err != nil {
		return
	}

	f.entries = make(map[string]*File)

	for k, v := range m {
		f.entries[k] = &File{
			Name:        k,
			Content:     v.Content,
			ContentType: v.ContentType,
		}
	}

	return nil
}

type File struct {
	Name        string
	Content     string `json:"content"`
	ContentType string `json:"content-type"`
}

