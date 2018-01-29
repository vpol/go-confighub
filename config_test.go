package confighub

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestClientGet(t *testing.T) {

	sample := []byte(`{
		"generatedOn": "01/22/2018 11:30:59",
		"account": "vpol",
		"repo": "api",
		"context": "production;auth;1",
    	"files": {
        	"test": {
            	"content": "adsfdfasadsfadfsdaf",
            	"content-type": "text/plain"
        	},
			"test1": {
            	"content": "11111111",
            	"content-type": "text/html"
        	}
    	},
    	"properties": {
			 "mapKey": {
            	"type": "Map",
            	"val": {
					"key1": "value1",
					"key2": "value2"
            	}
        	},
			"jsonKey": {
				"type": "JSON",
				"val": "{\"test\": 123}"
			},
			"stringKey": {
				"val": "1q2w3e"
			},
			"longKey": {
				"type": "Long",
				"val": "1212321312313"
			},
			"boolKey": {
				"type": "Boolean",
				"val": "true"
			},
			"intKey": {
				"type": "Integer",
				"val": "111"
			},
			"fileRefKey": {
				"type": "FileRef",
				"val": "test"
			}
    	}
	}`)

	var cfg ConfigHubClient

	assert.Nil(t, json.Unmarshal(sample, &cfg))

	assert.EqualValues(t, "vpol", cfg.Account)

	{
		_, err := cfg.GetProperty("unknownKey").String()
		assert.NotNil(t, err)
	}

	{
		val, err := cfg.GetProperty("stringKey").String()
		assert.Nil(t, err)
		assert.EqualValues(t, "1q2w3e", val)
	}

	{
		val, err := cfg.GetProperty("mapKey").StringMap()
		assert.Nil(t, err)
		assert.EqualValues(t, map[string]string{
			"key1": "value1",
			"key2": "value2",
		}, val)
	}

	{
		val, err := cfg.GetProperty("boolKey").Boolean()
		assert.Nil(t, err)
		assert.True(t, val)
	}

	{
		val, err := cfg.GetProperty("jsonKey").Json()
		assert.Nil(t, err)
		assert.EqualValues(t, map[string]interface{}{"test": float64(123)}, val)
	}

	{
		val, err := cfg.GetProperty("longKey").Long()
		assert.Nil(t, err)
		assert.EqualValues(t, 1212321312313, val)
	}

	{
		val, err := cfg.GetProperty("intKey").Integer()
		assert.Nil(t, err)
		assert.EqualValues(t, 111, val)
	}

	{
		val, err := cfg.GetProperty("fileRefKey").File()
		assert.Nil(t, err)
		assert.EqualValues(t, "adsfdfasadsfadfsdaf", val.Content)
		assert.EqualValues(t, "text/plain", val.ContentType)
	}
}
