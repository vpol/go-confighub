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
        	}
    	},
    	"properties": {
			"jsonKey": {
				"type": "JSON",
				"val": "{\"test\": 123}"
			},
			"token": {
				"val": "1q2w3e"
			},
			"longValue": {
				"type": "Long",
				"val": "1212321312313"
			},
			"cachedb": {
				"val": "default"
			},
			"login.method.token": {
				"type": "Boolean",
				"val": "true"
			},
			"intKey": {
				"type": "Integer",
				"val": "111"
			},
			"test": {
				"type": "FileRef",
				"val": "test"
			}
    	}
	}`)

	var cfg ConfigHubClient

	assert.Nil(t, json.Unmarshal(sample, &cfg))

	t.Log(cfg.Account)

	assert.NotNil(t, cfg.GetProperty("login.method.token"))

	{
		val, err := cfg.GetProperty("login.method.token").Boolean()
		assert.Nil(t, err)
		assert.True(t, val)
	}

	{
		val, err := cfg.GetProperty("jsonKey").Json()
		assert.Nil(t, err)
		assert.EqualValues(t, map[string]interface{}{"test": float64(123)}, val)
	}

	{
		val, err := cfg.GetProperty("longValue").Long()
		assert.Nil(t, err)
		assert.EqualValues(t, 1212321312313, val)
	}
}
