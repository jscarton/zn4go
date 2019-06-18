package lambdas

type ZnLambdaResponse struct {
	StatusCode int `json: "status"`
	Data map[string]interface{} `json: "data"`
	DeveloperMessage string `json: "developer"`
	UserMessage string `json: "message"`
}