package confighub

import (
	"encoding/json"
)

type Properties struct {
	cfg     *ConfigHubClient
	entries map[string]*Property
}

type Property struct {
	cfg      *ConfigHubClient
	key      string
	value    Value
	updateCB func(value Value)
}

func (p *Property) UnmarshalJSON(data []byte) (err error) {

	var m = make(map[string]interface{})

	err = json.Unmarshal(data, &m)

	if err != nil {
		return
	}

	valI, ok := m["val"]
	if !ok {
		err = NoValueErr
		return
	}

	if t, ok := m["type"]; !ok {

		valS := valI.(string)

		p.value = &String{
			Value: valS,
		}
	} else {

		switch t {
		case "Boolean":
			p.value = &Boolean{}
		case "Double":
			p.value = &Double{}
		case "Long":
			p.value = &Long{}
		case "Integer":
			p.value = &Integer{}
		case "Float":
			p.value = &Float{}
		case "JSON":
			p.value = &Json{}
		case "Map":
			p.value = &StringMap{}
		case "FileRef":
			p.value = &FileRef{
				files: p.cfg.Files,
			}
		default:
			err = WrongValueErr
			return
		}

		err = p.value.Parse(valI)
	}

	return nil
}
