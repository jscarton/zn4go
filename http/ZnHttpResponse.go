package http

import "encoding/json"

type ZnHttpResponseSingleRecord struct {
	Status int `json: "status"`
	Code int `json: "code"`
	TotalCount int `json: "totalCount"`
	Limit int `json: "limit"`
	Offset int `json: "offset"`
	Data map[string]interface{} `json: "data"`
	UserMessage string `json: "userMessage"`
	DeveloperMessage string `json: "developerMessage"`
	ValidationErrors map[string]interface{} `json: "validationErrors"`
	Error string `json: "error"`
	Error_Description string `json: "error_description"`
}

type ZnHttpResponse struct {
	Status int `json: "status"`
	Code int `json: "code"`
	TotalCount int `json: "totalCount"`
	Limit int `json: "limit"`
	Offset int `json: "offset"`
	Data []map[string]interface{} `json: "data"`
	UserMessage string `json: "userMessage"`
	DeveloperMessage string `json: "developerMessage"`
	ValidationErrors map[string]interface{} `json: "validationErrors"`
	Error string `json: "error"`
	Error_Description string `json: "error_description"`
}

func FormatResponse(raw_response []byte) (ZnHttpResponse, error) {
	response := ZnHttpResponse{}
	err := json.Unmarshal(raw_response, &response)
	if err != nil {
		resp := ZnHttpResponseSingleRecord{}
		err = json.Unmarshal(raw_response, &resp)
		if (err == nil) {
			response.Status = resp.Status
			response.Code = resp.Code
			response.TotalCount = resp.TotalCount
			response.Limit = resp.Limit
			response.Offset = resp.Offset
			response.Data = []map[string]interface{}{resp.Data}
			response.UserMessage = resp.UserMessage
			response.DeveloperMessage = resp.DeveloperMessage
			response.ValidationErrors = resp.ValidationErrors
			response.Error = resp.Error
			response.Error_Description = resp.Error_Description
		} else {
			return response,err	
		}
	}
	return response,err
}