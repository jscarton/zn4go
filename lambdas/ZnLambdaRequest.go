package lambdas

import (
	"github.com/jscarton/zn4go/http"
	"strings"
)

type ZnLambdaRequest struct {
	Workspace map[string]string `json: "workspace"`
	Authorization string `json: "authorization"`
	Params map[string]interface{} `json: "params"`
}

func (r *ZnLambdaRequest) GetHttpClient() (http.ZnHttp,error) {
	c,err := http.NewZnHttp(r.Authorization, nil)
	return c,err
}

func (r *ZnLambdaRequest) GetWorkspaceId() string {
	return r.Workspace["id"]
}

func (r *ZnLambdaRequest) GetParams() map[string]interface{} {
	return r.Params
}

func (r *ZnLambdaRequest) GetParam(path string) interface{} {
	var string_path = strings.Split(path,".")
	mappedObj := r.Params
	obj,is := mappedObj[string_path[0]];
	if is {
		if (len(string_path) > 1) {
			return r.GetParamFrom(string_path[1:],obj)
		} else {
			return obj
		}
	} else {
		return nil
	}
}

func (r *ZnLambdaRequest) GetParamFrom(path []string, obj interface{} ) interface{} {
	mappedObj := obj.(map[string]interface{})
	innerObj,is := mappedObj[path[0]];
	if is {
		if (len(path) > 1) {
			return r.GetParamFrom(path[1:],innerObj)
		} else {
			return innerObj
		}
	} else {
		return nil
	}
}