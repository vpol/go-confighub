package confighub

import (
	"time"
	"encoding/json"
	"sync"
	"fmt"
	"net/http"
)

type ConfigProvider interface {
	GetProperty(string) Value
	GetFile(string) *File
	AddWatch(string, WatchFunc)
}

type ConfigHubClient struct {
	mu             sync.Mutex
	options        *ConfigOptions
	watches        map[string][]WatchFunc
	Account        string          `json:"account"`
	GenerationTime time.Time       `json:"generated_on"`
	Repo           string          `json:"repo"`
	Context        string          `json:"context"`
	PropertiesB    json.RawMessage `json:"properties,omitempty"`
	Properties     *Properties     `json:"-"`
	FilesB         json.RawMessage `json:"files,omitempty"`
	Files          *Files          `json:"-"`
	client         *http.Client
}

func NewConfigHubClient(options *ConfigOptions) *ConfigHubClient {
	return &ConfigHubClient{
		options: options,
		watches: make(map[string][]WatchFunc),
	}
}

func (c *ConfigHubClient) UnmarshalJSON(data []byte) (err error) {

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)

	if err != nil {
		return
	}

	if accI, ok := m["account"]; ok {
		err = json.Unmarshal(accI, &c.Account)
		if err != nil {
			return
		}
	}

	if genI, ok := m["generated_on"]; ok {
		// 01/21/2018 23:02:47
		c.GenerationTime, err = time.Parse("01/02/2006 15:04:05", string(genI))
		if err != nil {
			return
		}
	}

	if repoI, ok := m["repo"]; ok {
		err = json.Unmarshal(repoI, &c.Repo)
		if err != nil {
			return
		}
	}

	if ctxI, ok := m["context"]; ok {
		err = json.Unmarshal(ctxI, &c.Context)
		if err != nil {
			return
		}

	}

	c.Files = &Files{
		cfg:     c,
		entries: make(map[string]*File),
	}

	if fI, ok := m["files"]; ok {

		var m map[string]json.RawMessage

		err = json.Unmarshal(fI, &m)

		if err != nil {
			return
		}

		for k, v := range m {

			var file File

			if err = json.Unmarshal(v, &file); err != nil {
				return
			}

			file.Name = k
			c.Files.entries[k] = &file

		}

	}

	c.Properties = &Properties{
		cfg:     c,
		entries: make(map[string]*Property),
	}

	if pI, ok := m["properties"]; ok {

		var m map[string]json.RawMessage

		err = json.Unmarshal(pI, &m)

		if err != nil {
			return
		}

		for k, v := range m {

			var property = Property{
				cfg: c,
			}

			if err = json.Unmarshal(v, &property); err != nil {
				return
			}

			property.cfg = c
			property.key = k
			c.Properties.entries[k] = &property

		}
	}

	return nil
}

func (c *ConfigHubClient) GetProperty(key string) (v Value) {

	if p, ok := c.Properties.entries[key]; ok {
		return p.value
	}

	return &NilValue{}
}

func (c *ConfigHubClient) GetFile(name string) (f *File) {

	if c.Files == nil {
		return nil
	}

	if f, ok := c.Files.entries[name]; ok {
		return f
	}

	return nil
}

func (c *ConfigHubClient) fetch() (result interface{}, err error) {

	if c.client == nil {
		c.client = &http.Client{}
	}

	var req *http.Request

	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/rest/pull/%s/api", c.options.Host, c.options.Repo), nil)

	req.Header.Set("Context", c.options.Context)

	return c.client.Do(req)
}

type WatchFunc func(val Value) error

func (c *ConfigHubClient) AddWatch(key string, watch WatchFunc) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if watchers, ok := c.watches[key]; ok {
		watchers = append(watchers, watch)
	} else {
		watchers = []WatchFunc{watch}
	}
}

type ConfigOption func(*ConfigOptions)

func DefaultOptions() *ConfigOptions {
	return &ConfigOptions{
		Host: "127.0.0.1",
		Port: 8080,
		Token: "secret",
		Repo: "default",
		Context: "",
	}
}

type ConfigOptions struct {
	Host    string `json:"host"`
	Port    int64  `json:"port"`
	Token   string `json:"token"`
	Repo    string `json:"repo"`
	Context string `json:"context"`
}

func Host(host string) ConfigOption {
	return func(config *ConfigOptions) {
		config.Host = host
	}
}

func Port(port int64) ConfigOption {
	return func(config *ConfigOptions) {
		config.Port = port
	}
}

func Token(token string) ConfigOption {
	return func(config *ConfigOptions) {
		config.Token = token
	}
}

func Repo(repo string) ConfigOption {
	return func(config *ConfigOptions) {
		config.Repo = repo
	}
}

func Context(context string) ConfigOption {
	return func(config *ConfigOptions) {
		config.Context = context
	}
}
