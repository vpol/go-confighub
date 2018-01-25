package confighub

import (
	"encoding/json"
	"fmt"
)

type Properties struct {
	cfg      *ConfigHubClient
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

	valS := valI.(string)

	if t, ok := m["type"]; !ok {
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
		case "FileRef":
			p.value = &FileRef{
				files: p.cfg.Files,
			}
		default:
			fmt.Println(t)
			err = WrongValueErr
			return
		}

		err = p.value.Parse(valS)
	}

	return nil
}
