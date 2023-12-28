package handler

type SuccessResponse struct {
	Result interface{} `json:"result"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}

//
