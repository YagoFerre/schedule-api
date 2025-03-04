package rest_errors

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}
