package config

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
)

type ZnConfig struct {
	Version int64
	Data map[string]interface{}
}


func Init(path string) (ZnConfig, error) {
	data, err := ioutil.ReadFile(path)
	settings := ZnConfig{}
	now := time.Now()
    secs := now.Unix()
    settings.Version = secs
	// if can't read the file set the errString and return false
	if err != nil {
		return settings, err
	}
	err = json.Unmarshal(data,&settings.Data)
	if err != nil {
		return settings, err
	}
	return settings, nil
}

func (c *ZnConfig) Get(path string) interface{} {
	var string_path = strings.Split(path,".")
	mappedObj := c.Data
	obj,is := mappedObj[string_path[0]];
	if is {
		if (len(string_path) > 1) {
			return c.GetFrom(string_path[1:],obj)
		} else {
			return obj
		}
	} else {
		return nil
	}
}

func (c *ZnConfig) GetFrom(path []string, obj interface{} ) interface{} {
	mappedObj := obj.(map[string]interface{})
	innerObj,is := mappedObj[path[0]];
	if is {
		if (len(path) > 1) {
			return c.GetFrom(path[1:],innerObj)
		} else {
			return innerObj
		}
	} else {
		return nil
	}
}

func (c *ZnConfig) GetVersion() int64 {
	return c.Version;
}