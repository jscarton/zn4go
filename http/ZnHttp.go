package http

import (
	"net/url"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type ZnHttp struct {
	access_token string
	base_url string
	default_params map[string]string
}


func NewZnHttp(at string,dp map[string]string) (ZnHttp,error) {
	return newZnHttp(at, "https://api.zenginehq.com/v1",dp)
}

func newZnHttp(at string, bu string, dp map[string]string) (ZnHttp, error) {

	obj := ZnHttp{};

	if len(at) == 0 {
		return obj, errors.New("ZnHttp: invalid access_token")
	} else {
		obj.access_token = at
	}

	if len(bu) == 0 {
		return obj, errors.New("ZnHttp: invalid url.")
	} else {
		_, err := url.ParseRequestURI(bu)
		if err != nil {
			return obj, errors.New("ZnHttp: invalid url."+err.Error())
		} else {
			obj.base_url = bu
		}
	}
	
	obj.default_params = dp
	return obj, nil	 
}

func (c *ZnHttp) Get(path string, params map[string]string) (ZnHttpResponse,error) {
	requestUrl := c.base_url+path+"?access_token="+c.access_token
	response := ZnHttpResponse{}
	_, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		return response, errors.New("ZnHttp.Get: invalid url."+err.Error())
	}
	if len(c.default_params) > 0 || len(params) >0 {
		requestUrl+="?"
		for k,v:= range c.default_params {
			requestUrl+=k+"="+v
		}
		for k,v:= range params {
			requestUrl+=k+"="+v
		}
	}
	resp, err := http.Get(requestUrl)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response, err = FormatResponse(body)
	return response,err
}

func (c *ZnHttp) Post(path string, params map[string]string, data interface{}) (ZnHttpResponse, error){
	requestUrl := c.base_url+path+"?access_token="+c.access_token
	response := ZnHttpResponse{}
	_, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		return response, errors.New("ZnHttp.Post: invalid url."+err.Error())
	}
	if len(c.default_params) > 0 || len(params) >0 {
		requestUrl+="?"
		for k,v:= range c.default_params {
			requestUrl+=k+"="+v
		}
		for k,v:= range params {
			requestUrl+=k+"="+v
		}
	}
	bytesRepresentation, err := json.Marshal(data)
	if err != nil {
		return response, errors.New("ZnHttp.Post: invalid payload."+err.Error())
	}
	resp, err := http.Post(requestUrl,"application/json",bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response, err = FormatResponse(body)
	return response,err
}

func (c *ZnHttp) Put(path string, params map[string]string, data interface{}) (ZnHttpResponse, error){
	requestUrl := c.base_url+path+"?access_token="+c.access_token
	response := ZnHttpResponse{}
	_, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		return response, errors.New("ZnHttp.Put: invalid url."+err.Error())
	}
	if len(c.default_params) > 0 || len(params) >0 {
		requestUrl+="?"
		for k,v:= range c.default_params {
			requestUrl+=k+"="+v
		}
		for k,v:= range params {
			requestUrl+=k+"="+v
		}
	}
	bytesRepresentation, err := json.Marshal(data)
	if err != nil {
		return response, errors.New("ZnHttp.Put: invalid payload."+err.Error())
	}
	req, err := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return response, errors.New("ZnHttp.Put: invalid request."+err.Error())
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	Client := http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response, err = FormatResponse(body)
	return response,err
}

func (c *ZnHttp) Delete(path string, params map[string]string) (ZnHttpResponse, error){
	requestUrl := c.base_url+path+"?access_token="+c.access_token
	response := ZnHttpResponse{}
	_, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		return response, errors.New("ZnHttp.Delete: invalid url."+err.Error())
	}
	if len(c.default_params) > 0 || len(params) >0 {
		requestUrl+="?"
		for k,v:= range c.default_params {
			requestUrl+=k+"="+v
		}
		for k,v:= range params {
			requestUrl+=k+"="+v
		}
	}
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return response, errors.New("ZnHttp.Delete: invalid request."+err.Error())
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	Client := http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response, err = FormatResponse(body)
	return response,err
}